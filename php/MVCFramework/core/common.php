<?php
//公共函数库
namespace core;

//打印
function dump(...$data)
{
    foreach ($data as $value) {
        $result = var_export($value,true);
//        自定义显示样式
        $style = 'border:1px solid #ccc;border-radius:5px;';
        $style .= 'padding:8px;background:#efefef;';
        printf('<pre style="%s">%s</pre>', $style,$result);
    }
}
