<?php

namespace app\admin\controller;

class CustomerController extends BaseController
{
    public function index(int $page = 1)
    {
        $table = 'customer';
        $limit = CONFIG['database']['default_limit'];
        $status =0;
        $data = $this->getPageData($table, $page, $limit, $status);
        // app/view/customer/index.php
        $this->view->render(null, $data);
    }
}