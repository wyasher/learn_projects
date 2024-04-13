// The Swift Programming Language
// https://docs.swift.org/swift-book
let expr: SExpr = "(cond ((atom (quote A)) (quote B)) ((quote true) (quote C)))"

print(expr)
print(expr.eval()!)  //B
var exit = false
while(!exit){
    print(">>>", terminator:" ")
    let input = readLine(strippingNewline: true)
    exit = (input=="exit") ? true : false

    if !exit {
        let e = SExpr.read(input!)
        print(e.eval()!)
    }
}