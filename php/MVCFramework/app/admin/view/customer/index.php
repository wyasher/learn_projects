<?php include __DIR__ . '/../public/header.php' ?>
<?php include __DIR__ . '/../public/left.php' ?>
<style>
    <?php include __DIR__ .'/../static/css/show.css'?>
</style>

<div class="content">
    <?php if (count($result) == 0)  : ?>
        <h2 class="title">没有满足条件的记录</h2>
    <?php else : ?>
        <table>
            <caption>客户信息列表</caption>
            <thead>
            <tr>
                <th>ID</th>
                <th>姓名</th>
                <th>性别</th>
                <th>生日</th>
                <th>地址</th>
                <th>邮箱</th>
                <th>手机号</th>
                <th>更新时间</th>
                <th>操作</th>
            </tr>
            </thead>

            <tbody>
            <?php foreach ($result as $customer): ?>
                <tr>
                    <td><?= $customer['id'] ?>
                    </td>
                    <td><?= $customer['cname'] ?>
                    </td>
                    <td><?= $customer['sex'] ? '女' : '男' ?>
                    </td>
                    <td><?= $customer['birthday'] ?>
                    </td>
                    <td><?= $customer['address'] ?>
                    </td>
                    <td><?= $customer['email'] ?>
                    </td>
                    <td><?= $customer['mobile'] ?>
                    </td>
                    <td><?= $customer['updated_time'] ?>
                    </td>
                    <td>
                        <button class="edit">编辑</button>
                        <button class="del">删除</button>
                    </td>
                </tr>
            <?php endforeach ?>
            </tbody>
        </table>

        <!-- 分页条 -->
        <nav class="paginate">

            <?php foreach ($paginate as $p) : ?>

                <?php if (is_null($p)) : ?>
                    <span style="text-align=center">...</span>
                <?php else : ?>

                    <?php $active = $p == $page ? 'active' : null ?>
                    <?php $url = '/admin.php/customer/index/' . $p ?>
                    <a href="<?= $url ?>"
                       class="<?= $active ?>"><?= $p ?></a>
                <?php endif ?>

            <?php endforeach ?>

        </nav>
    <?php endif ?>
</div>

