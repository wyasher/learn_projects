<?php
//数据库操作类
namespace core;

use PDO;

class Db{
    protected static PDO $db;
    public static function connect(string $dsn, string $username, string $password){
        self::$db = new PDO($dsn,$username,$password);
    }
    public static function __callStatic(string $name,array $arguments)
    {
        $dsn = CONFIG['database']['type'] . ':host=' . CONFIG['database']['host'] . ';dbname=' . CONFIG['database']['dbname'];
        $username = CONFIG['database']['username'];
        $password = CONFIG['database']['password'];
        static::connect($dsn,$username,$password);
        $query = new Query(self::$db);
        return call_user_func_array([$query,$name],$arguments);
    }

}