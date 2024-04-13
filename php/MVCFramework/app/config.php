<?php

// 应用默认配置文件
return [
    'app' => [
        'debug' => true,
        'default_controller' => 'Index',
        'default_action' => 'customer',
        'default_view_suffix' => 'php',
    ],
    'database' => [
        'type' => 'mysql',
        'host' => '127.0.0.1',
        'port' => 3306,
        'dbname' => 'admin',
        'charset' => 'utf8',
        'username' => 'root',
        'password' => 'root',
        'default_limit' => 10
    ]
];