<?php

// 框架应用基础类，启动框架
namespace core;
class App
{
    public static function run()
    {
        // 1 启动回话
        session_start();
        // 2 加载公共库函数
        require __DIR__ . DIRECTORY_SEPARATOR . 'common.php';
        // 3 设置常量
        self::setConstants();
        // 4 类自动加载
        spl_autoload_register([__CLASS__, 'autoloader']);
        // 5 路由解析
        [$controller,$action,$params] = Router::parse();
        // 6 实例化控制器
        // 获取视图实例
        $appPath = ROOT_APP_PATH . DIRECTORY_SEPARATOR . APP_NAME;
        $path = $appPath . DIRECTORY_SEPARATOR . 'view';
        $view = new View($path,$controller,$action);
        $controller ='app\\' . APP_NAME . '\\controller\\'   . ucfirst($controller) . 'Controller';
        $controller = new $controller($view);
        // 调用控制器方法
        echo call_user_func_array([$controller,$action],$params);
    }

    // 类自动加载
    private static function autoloader($class)
    {
        $file = str_replace('\\', DIRECTORY_SEPARATOR, $class) . '.php';
        // 2 判断类文件是否存在
        file_exists($file) ? require $file : die($file . '类文件不存在');
    }

    // 设置常量
    private static function setConstants()
    {
        define('CORE_PATH', __DIR__);
        define('ROOT_PATH', dirname(__DIR__));
        define('ROOT_APP_PATH', ROOT_PATH . DIRECTORY_SEPARATOR . 'app');

        // 默认配置
        $defaultConfig = ROOT_PATH . DIRECTORY_SEPARATOR . 'config.php';
        $appConfig = ROOT_APP_PATH . DIRECTORY_SEPARATOR . 'config.php';

        define('CONFIG', require file_exists($appConfig) ? $appConfig : $defaultConfig);
        // 设置调试
        ini_set('display_errors', CONFIG['app']['debug'] ? 'On' : 'Off');
    }
}