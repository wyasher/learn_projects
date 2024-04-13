# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: "Star Wars" }, { name: "Lord of the Rings" }])
#   Character.create(name: "Luke", movie: movies.first)
Product.delete_all
Product.create!(
  title: "Apple MacBook Pro 13英寸 M2 芯片",
  description: "Apple MacBook Pro 13英寸 M2 芯片(8核中央处理器 10核图形处理器) 16G 512G 深空灰 笔记本Z16S【定制机】",
  image_url: "https://img11.360buyimg.com/n7/jfs/t1/190610/17/25727/35590/62b27acdEd861c52d/bb992d328c539da8.jpg",
  price: 12999.00
)
Product.create!(
  title: "华为笔记本电脑MateBook 14 2023",
  description: "华为笔记本电脑MateBook 14 2023 13代酷睿版 i5 16G 512G 14英寸轻薄办公本/2K触控全面屏/手机互联 深空灰",
  image_url: "https://img12.360buyimg.com/n7/jfs/t1/146713/22/30645/61866/648bc2beF3c042622/0e406d23dc2a1a2a.jpg",
  price: 5499.00
)
Product.create!(
  title: "华为P50E手机",
  description: "华为P50E手机 鸿蒙操作系统 5000万超感光原色影像 支持66W快充 8GB+256GB可可茶金",
  image_url: "https://img13.360buyimg.com/n1/s450x450_jfs/t1/223935/39/25110/33953/648148cbF9b005737/db6535dd0dc3e99f.jpg",
  price: 3499.00
)