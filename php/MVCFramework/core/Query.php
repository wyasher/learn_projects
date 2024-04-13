<?php

namespace core;

use PDO;

class Query
{
    protected PDO $db;
    protected string $table;
    protected string $filed = "*";
    protected array $opt = [];

    public function __construct(PDO $db)
    {
        $this->db = $db;
    }


    public function where(string $where): self
    {
        $this->opt['where'] = " where $where";
        return $this;
    }

    public function order(string $filed, string $order = "DESC"): self
    {
        $this->opt['order'] = " order by $filed $order";
        return $this;
    }

    public function page(int $currentPage = 1, int $pageSize = 10): self
    {
        $this->opt['page'] = " limit " . ($currentPage - 1) * $pageSize . ",$pageSize";
        return $this;
    }

    public function table(string $table): self
    {
        $this->table = $table;
        return $this;
    }

    public function filed(string $filed = "*"): self
    {
        $this->filed = $filed;
        return $this;
    }

    public function select(): array
    {
        $sql = "select $this->filed from $this->table";

        if (isset($this->opt['where'])) {
            $sql .= $this->opt['where'];
        }
        if (isset($this->opt['order'])) {
            $sql .= $this->opt['order'];
        }
        if (isset($this->opt['page'])) {
            $sql .= $this->opt['page'];
        }
        $this->opt = [];
        $stmt = $this->db->prepare($sql);
        $stmt->execute();
        return $stmt->fetchAll(PDO::FETCH_ASSOC);
    }

    public function find()
    {
        $sql = "select $this->filed from $this->table";

        if (isset($this->opt['where'])) {
            $sql .= $this->opt['where'];
        }
        $this->opt = [];
        $stmt = $this->db->prepare($sql);
        $stmt->execute();
        return $stmt->fetch(PDO::FETCH_ASSOC);
    }

    public function insert(array $data): int
    {
        $str = '';
        foreach ($data as $key => $value) {
            $str .= $key . ' = ' . "'$value'" . ',';
        }
        $str = rtrim($str, ',');
        $sql = "insert  $this->table set $str";
        $stmt = $this->db->prepare($sql);
        $stmt->execute();
        return $stmt->rowCount();
    }

    public function update(array $data): int
    {
        $str = '';
        foreach ($data as $key => $value) {
            $str .= $key . ' = ' . "'$value'" . ',';
        }
        $sql = 'UPDATE ' . $this->table . ' SET ' . rtrim($str, ', ');
        $sql .= $this->opt['where'] ?? die('禁止无条件更新');
        $stmt = $this->db->prepare($sql);
        $stmt->execute();
        return $stmt->rowCount();
    }

    public function delete(): int
    {
        $sql = 'DELETE FROM ' . $this->table;
        $sql .= $this->opt['where'] ?? die('禁止无条件删除');
        $stmt = $this->db->prepare($sql);
        $stmt->execute();
        return $stmt->rowCount();
    }
}