//
// Created by 王洋 on 2023/6/2.
//

import Foundation

public enum SExpr {
    case Atom(String)
    case List([SExpr])

    public func eval(with locals: [SExpr]? = nil, for values: [SExpr]? = nil) -> SExpr?{
        var node = self

        switch node {
        case .Atom:
            return evaluateVariable(node, with:locals, for:values)
        case var .List(elements):
            var skip = false

            if elements.count > 1, case let .Atom(value) = elements[0] {
                skip = Builtins.mustSkip(value)
            }

            // Evaluate all subexpressions
            if !skip {
                elements = elements.map{
                    return $0.eval(with:locals, for:values)!
                }
            }
            node = .List(elements)

            // Obtain a a reference to the function represented by the first atom and apply it, local definitions shadow global ones
            if elements.count > 0, case let .Atom(value) = elements[0], let f = localContext[value] ?? defaultEnvironment[value] {
                let r = f(node,locals,values)
                return r
            }

            return node
        }
    }

    private func evaluateVariable(_ v: SExpr, with locals: [SExpr]?, for values: [SExpr]?) -> SExpr {
        guard let locals = locals, let values = values else {return v}

        if locals.contains(v) {
            // The current atom is a variable, replace it with its value
            return values[locals.index(of: v)!]
        }else{
            // Not a variable, just return it
            return v
        }
    }
}

// 判断表达式是否相等
extension SExpr: Equatable {
    public static func ==(lhs: SExpr, rhs: SExpr) -> Bool {
        switch (lhs, rhs) {
        case let (.Atom(l), .Atom(r)):
            return l == r
        case let (.List(l), .List(r)):
            guard l.count == r.count else {
                return false
            }
            // 判断List中所有是否相同
            for (index, element) in l.enumerated() {
                if element != r[index] {
                    return false
                }
            }
            return true
        default:
            return false
        }
    }
}

extension SExpr: CustomStringConvertible {
    public var description: String {
        switch self {
        case let .Atom(value):
            return "\(value)"
        case let .List(exprList):
            var res = "("
            for expr in exprList {
                res += "\(expr)"
            }
            res += ")"
            return res
        }
    }
}

//字符串字面量获得一个表达式
extension SExpr: ExpressibleByStringLiteral,
        ExpressibleByUnicodeScalarLiteral,
        ExpressibleByExtendedGraphemeClusterLiteral {
    public init(stringLiteral value: String) {
        self = SExpr.read(value)
    }

    public init(extendedGraphemeClusterLiteral value: String) {
        self = SExpr.read(value)
    }

    public init(unicodeScalarLiteral value: String) {
        self = SExpr.read(value)
    }

}

// 词法分析
extension SExpr {
    /**
    读取 LISP 字符串表达式并将其转换为具有层级结构的 S-Expression
     - Parameter expr: LISP 字符串表达式
     - Returns: S-Expression
     */
    public static func read(_ expr: String) -> SExpr {
        enum Token {
            case pOpen, pClose, textBlock(String)
        }

        /**
         将字符串转换为此法单元
         - Parameter expr: LISP 字符串表达式
         */
        func tokenize(_ expr: String) -> [Token] {
            var res = [Token]()
            var tempText = ""
            for character in expr {
                switch character {
                case "(":
                    if !tempText.isEmpty {
                        res.append(.textBlock(tempText))
                        tempText = ""
                    }
                    res.append(.pOpen)
                case ")":
                    if !tempText.isEmpty {
                        res.append(.textBlock(tempText))
                        tempText = ""
                    }
                    res.append(.pClose)
                case " ":
                    if !tempText.isEmpty {
                        res.append(.textBlock(tempText))
                        tempText = ""
                    }
                default:
                    tempText.append(character)
                }
            }
            return res
        }

        func appendTo(list: SExpr?, node: SExpr) -> SExpr {
            var list = list
            if list != nil, case var .List(elements) = list! {
                elements.append(node)
                list = .List(elements)
            } else {
                list = node
            }
            return list!
        }

        // 语法分析
        func parse(_ tokens: [Token], node: SExpr? = nil) -> (remaining: [Token], subExpr: SExpr) {
            var tokens = tokens
            var node = node
            var i = 0

            repeat {
                let token = tokens[i]
                switch (token) {
                case .pOpen:
                    let (tr, n) = parse(Array(tokens[(i + 1)..<tokens.count]), node: .List([]))
                    (tokens, i) = (tr, 0)
                    node = appendTo(list: node, node: n)
                    if tokens.count != 0 {
                        continue
                    } else {
                        break
                    }
                case .pClose:
                    // closer
                    return (Array(tokens[(i + 1)..<tokens.count]), node!)
                case let .textBlock(value):
                    node = appendTo(list: node, node: .Atom(value))
                }
                i += 1

            } while tokens.count > 0

            return ([], node!)
        }


        let tokens = tokenize(expr)
        let res = parse(tokens)
        return res.subExpr


    }
}

fileprivate enum Builtins: String {
    case quote, car, cdr, cons, equal, atom, cond, lambda, defun, list, println, eval

    public static func mustSkip(_ atom: String) -> Bool {
        (atom == Builtins.quote.rawValue) ||
                (atom == Builtins.cond.rawValue) ||
                (atom == Builtins.defun.rawValue) ||
                (atom == Builtins.lambda.rawValue)
    }
}

public var localContext = [String: (SExpr, [SExpr]?, [SExpr]?) -> SExpr]()

// 默认执行环境
private var defaultEnvironment: [String: (SExpr, [SExpr]?, [SExpr]?)->SExpr] = {

    var env = [String: (SExpr, [SExpr]?, [SExpr]?)->SExpr]()
    env[Builtins.quote.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 2 else {return .List([])}
        return parameters[1]
    }
    env[Builtins.car.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 2 else {return .List([])}
        guard case let .List(elements) = parameters[1], elements.count > 0 else {return .List([])}

        return elements.first!
    }
    env[Builtins.cdr.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 2 else {return .List([])}

        guard case let .List(elements) = parameters[1], elements.count > 1 else {return .List([])}

        return .List(Array(elements.dropFirst(1)))
    }
    env[Builtins.cons.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 3 else {return .List([])}

        guard case .List(let elRight) = parameters[2] else {return .List([])}

        switch parameters[1].eval(with: locals,for: values)!{
        case let .Atom(p):
            return .List([.Atom(p)]+elRight)
        default:
            return .List([])
        }
    }
    env[Builtins.equal.rawValue] = {params,locals,values in
        guard case let .List(elements) = params, elements.count == 3 else {return .List([])}

        var me = env[Builtins.equal.rawValue]!

        switch (elements[1].eval(with: locals,for: values)!,elements[2].eval(with: locals,for: values)!) {
        case (.Atom(let elLeft),.Atom(let elRight)):
            return elLeft == elRight ? .Atom("true") : .List([])
        case (.List(let elLeft),.List(let elRight)):
            guard elLeft.count == elRight.count else {return .List([])}
            for (idx,el) in elLeft.enumerated() {
                let testeq:[SExpr] = [.Atom("Equal"),el,elRight[idx]]
                if me(.List(testeq),locals,values) != SExpr.Atom("true") {
                    return .List([])
                }
            }
            return .Atom("true")
        default:
            return .List([])
        }
    }
    env[Builtins.atom.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 2 else {return .List([])}

        switch parameters[1].eval(with: locals,for: values)! {
        case .Atom:
            return .Atom("true")
        default:
            return .List([])
        }
    }
    env[Builtins.cond.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count > 1 else {return .List([])}

        for el in parameters.dropFirst(1) {
            guard case let .List(c) = el, c.count == 2 else {return .List([])}

            if c[0].eval(with: locals,for: values) != .List([]) {
                let res = c[1].eval(with: locals,for: values)
                return res!
            }
        }
        return .List([])
    }
    env[Builtins.defun.rawValue] =  { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 4 else {return .List([])}

        guard case let .Atom(lname) = parameters[1] else {return .List([])}
        guard case let .List(vars) = parameters[2] else {return .List([])}

        let lambda = parameters[3]

        let f: (SExpr, [SExpr]?, [SExpr]?)->SExpr = { params,locals,values in
            guard case var .List(p) = params else {return .List([])}
            p = Array(p.dropFirst(1))

            // Replace parameters in the lambda with values
            if let result = lambda.eval(with:vars, for:p){
                return result
            }else{
                return .List([])
            }
        }

        localContext[lname] = f
        return .List([])
    }
    env[Builtins.lambda.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 3 else {return .List([])}

        guard case let .List(vars) = parameters[1] else {return .List([])}
        let lambda = parameters[2]
        //Assign a name for this temporary closure
        let fname = "TMP$"+String(arc4random_uniform(UInt32.max))

        let f: (SExpr, [SExpr]?, [SExpr]?)->SExpr = { params,locals,values in
            guard case var .List(p) = params else {return .List([])}
            p = Array(p.dropFirst(1))
            //Remove temporary closure
            localContext[fname] = nil

            // Replace parameters in the lambda with values
            if let result = lambda.eval(with:vars, for:p){
                return result
            }else{
                return .List([])
            }
        }

        localContext[fname] = f
        return .Atom(fname)
    }
    //List implemented as a classic builtin instead of a series of cons
    env[Builtins.list.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count > 1 else {return .List([])}
        var res: [SExpr] = []

        for el in parameters.dropFirst(1) {
            switch el {
            case .Atom:
                res.append(el)
            case let .List(els):
                res.append(contentsOf: els)
            }
        }
        return .List(res)
    }
    env[Builtins.println.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count > 1 else {return .List([])}

        print(parameters[1].eval(with: locals,for: values)!)
        return .List([])
    }
    env[Builtins.eval.rawValue] = { params,locals,values in
        guard case let .List(parameters) = params, parameters.count == 2 else {return .List([])}

        return parameters[1].eval(with: locals,for: values)!
    }

    return env
}()