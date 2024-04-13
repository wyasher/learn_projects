<?php
namespace core;
// 视图类基类


class Controller{
    protected View $view;

    public function __construct(View $view){
        $this->view = $view;
    }
}