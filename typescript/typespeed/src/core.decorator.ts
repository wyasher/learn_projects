import "reflect-metadata"
import LogFactory from "./factory/log-factory";
import * as cron from "cron";

const resourceObjects = new Map<string, object>();
const beanMapper = new Map<string, any>();
const objectMapper = new Map<string, any>();


function component(constructorFunction: { new(): any; name: string; }) {
    objectMapper.set(constructorFunction.name, new constructorFunction());
}

function getComponent(constructorFunction: { new(): any; name: string; }) {
    return objectMapper.get(constructorFunction.name)
}

/**
 *    example
 *    @bean
 *     public createLog(): LogFactory {
 *         return new CustomLog();
 *     }
 */
function bean(target: any, propertyKey: string) {
    let returnType = Reflect.getMetadata("design:returntype", target, propertyKey);
    beanMapper.set(returnType.name, {
        "target": target,
        "propertyKey": propertyKey,
        "factory": target[propertyKey]()
    });
}

function getBean(mappingClass: Function): any {
    const bean = beanMapper.get(mappingClass.name);
    return bean["factory"];
}

function autoware(target: any, propertyKey: string): void {
    const type = Reflect.getMetadata("design:type", target, propertyKey);
    Object.defineProperty(target, propertyKey, {
        get: () => {
            const targetObject = beanMapper.get(type.name);
            if (targetObject === undefined) {
                const resourceKey = [target.constructor.name, propertyKey, type.name].toString();
                if (!resourceObjects[resourceKey]) {
                    resourceObjects[resourceKey] = new type();
                }
                return resourceObjects[resourceKey];
            }
            return targetObject["factory"];
        }
    });
}

function resource(...args: any[]): any {
    return (target: any, propertyKey: string) => {
        const type = Reflect.getMetadata("design:type", target, propertyKey);
        Object.defineProperty(target, propertyKey, {
            get: () => {
                const resourceKey = [target.constructor.name, propertyKey, type.name].toString();
                if (!resourceObjects[resourceKey]) {
                    resourceObjects[resourceKey] = new type(...args);
                }
                return resourceObjects[resourceKey];
            }
        });
    }
}

function log(message?: any, ...optionalParams: any[]) {
    const logObject = beanMapper.get(LogFactory.name);
    if (logObject) {
        logObject["factory"].log(message, ...optionalParams);
    } else {
        console.log(message, ...optionalParams);
    }
}

function error(message?: any, ...optionalParams: any[]) {
    const logObject = beanMapper.get(LogFactory.name);
    if (logObject) {
        logObject["factory"].error(message, ...optionalParams);
    } else {
        console.error(message, ...optionalParams);
    }
}

function schedule(cornTime:string | Date){
    return (target:any,propertyKey:string) => {
        new cron.CronJob(cornTime,target[propertyKey]).start();
    }
}

export { component, bean, resource, log, error, autoware, getBean, getComponent, schedule };
