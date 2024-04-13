<?php
namespace core;

// 获取当前模块名称/应用
define("APP_NAME",basename(__FILE__,".php"));
// 加载MVC核心
require __DIR__ . '/core/App.php';
// 启动框架
App::run();