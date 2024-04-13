<?php

namespace core;
class View
{
    protected string $defaultPath;
    protected string $controller;
    protected string $action;
    protected array $data = [];

    public function __construct(string $path, string $controller, string $action)
    {
        $this->defaultPath = $path;
        $this->controller = $controller;
        $this->action = $action;
    }
    // 为视图文件中的变量赋值
    // $key 变量名
    // $value 变量值
    public function assign($key, $value)
    {
        // 如果是数组，就合并
        if (is_array($value)){
            $this->data = array_merge($this->data,$value);
            return;
        }
        $this->data[$key] = $value;
    }
    // 渲染视图
    // $path 视图文件路径
    // $data 视图文件中需要的数据
    public function render(string $path = null, array $data = null)
    {
        $this->data = $data ? array_merge($this->data, $data) : $this->data;
        $file = $path ?? $this->controller . DIRECTORY_SEPARATOR . $this->action;
        $file = $file . '.' . CONFIG['app']['default_view_suffix'];
        $file = $this->defaultPath . DIRECTORY_SEPARATOR . $file;
        extract($this->data);
        file_exists($file) ? include $file : die($file . '视图文件不存在');
    }


}