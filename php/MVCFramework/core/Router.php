<?php
namespace core;

class Router{
    public static function parse():array{

        $controller = CONFIG['app']['default_controller'];
        $action = CONFIG['app']['default_action'];
        $params = [];
        if (array_key_exists('PATH_INFO',$_SERVER) && $_SERVER['PATH_INFO'] != '/') {
            $pathInfo = explode('/',trim($_SERVER['PATH_INFO'],'/'));
            // 获取控制器
            if (count($pathInfo) >= 2){
                $controller = array_shift($pathInfo);
                $action = array_shift($pathInfo);
                $params = $pathInfo;
            }else if (count($pathInfo) == 1){
                $controller = array_shift($pathInfo);
            }
        }

        // 返回控制器 方法 参数
        return [
            $controller,
            $action,
            $params
        ];
    }
}