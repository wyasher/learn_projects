<?php

namespace app\admin\controller;

use core\Controller;
use core\Db;
use function core\dump;

class BaseController extends Controller
{
    /**
     * 获取分页数据
     * @param string $table 表
     * @param int $page 当前页
     * @param int $limit 限制条数
     * @param string $state 状态
     * @return array 返回分页数据
     */
    protected function getPageData(string $table, int $page, int $limit, string $state): array
    {
        $result = Db::table($table)->page($page)
            ->where('status in (' . $state . ')')
            ->order('updated_time')
            ->select();
        $nums = Db::table($table)
            ->where('status in (' . $state . ')')
            ->select();

        $total = count($nums);
        $pages = intval(ceil($total / $limit));
        // 4. 生成分页码
        $paginate = $this->createPages($page, $pages);
        // dump($paginate);

        // 5. 返回分页视图中的所有数据
        return [
            // 分页数据
            'result' => $result,
            // 总页数
            'pages' => $pages,
            // 总记录数
            'total' => $total,
            // 分页数组
            'paginate' => $paginate,
            // 当前页
            'page' => $page
        ];
    }

    /**
     * 生成分页信息
     */
    private function createPages(int $page, int $pages): array
    {
        // 当前是第8页, 共计20页
        // [1,  ...    6, 7, 8, 9, 10,  ....    20]
        // 当前是第10页, 共计20页
        // [1,  ...    8, 9, 10, 11, 12,  ....    20]

        // 1. 生成与总页数长度相同的递增的整数数组
        $pageArr = range(1, $pages);

        // 2. 只需要当前和前后二页, 其它页码用 false/null 来标记

        $paginate =  array_map(function ($p) use ($page, $pages) {
            return   ($p == 1 || $p == $pages || abs($page-$p) <=2) ? $p : null;
        }, $pageArr);

        // dump($paginate);
        // 去重, 替换
        $before = array_unique(array_slice($paginate, 0, $page));
        $after = array_unique(array_slice($paginate, $page));


        // 用解构进行合并
        return [...$before, ...$after];
    }

    protected function isLogged(): bool
    {
        return false;
    }


}