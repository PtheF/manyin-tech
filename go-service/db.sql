
/* user */
/**
  email 和 phone 原本是 text，unique，但是报错了，
  查了一下是因为 unique 就需要索引，但是对于 text 只能索引前 n 位
  所以这里改成了 char 和 varchar
 */
create table user(
    uuid char(50) primary key,
    email varchar(100) unique,
    phone char(16) unique,
    nike_name char(50) not null,
    password char(65) not null,
    salt char(65) not null,
    reg_date timestamp not null
)charset=utf8mb4;

create table address (
    uuid char(50),
    id char(50) not null,
    name char(20) not null,
    addressee char(30) not null,
    address text(255) not null,
    phone char(20) not null,
    is_default tinyint(1) not null,
    ava tinyint(1) not null ,
    foreign key (uuid) references user(uuid),
    check(ava = 1),
    check(is_default = 0 or is_default = 1)
)charset=utf8mb4;

create table enterprise_info (
    user_id char(50),
    full_name varchar(50) not null,
    social_code varchar(60) unique,
    registry_money int(64),
    enter_address varchar(50),
    operator_name char(20),
    operator_email varchar(100),
    operator_phone char(16),
    create_date timestamp not null,
    status tinyint comment '1:审核中, 2:审核通过, 3:不通过',
    primary key(user_id),
    foreign key (user_id) references user(uuid)
)charset=utf8mb4;

create table product_types (
    prod_t_id tinyint primary key auto_increment,
    t_name varchar(50) not null,
    sub_text varchar(50) not null,
    description text not null
)charset=utf8mb4;

insert into product_types (t_name, sub_text, description)
values
    ('装配式建筑', '装配式建筑的简单说明', '装配式建筑的详细描述'),
    ('再生材料', '再生材料的简单说明', '再生材料的详细描述'),
    ('再生骨料', '再生骨料的简单说明', '再生骨料的详细描述'),
    ('再生工艺品', '再生工艺品的简单说明', '再生工艺品的详细描述');

insert into enterprise_info
(

 user_id,
 full_name,
 social_code,
 registry_money,
 enter_address,
 operator_name,
 operator_email,
 operator_phone,
 create_date,
 status
)values (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, 1
);

update product_types set description = '此产品可用于装配式建筑的一体化建筑设计和施工，在工厂进行加工后，直接运送至工地进行装配，建筑时无需搭建外架、手脚板等，可以有效节约人力和物力。装配式构配件尚可多次使用，减少建筑垃圾的产生。' where prod_t_id = 1;
update product_types set description = '蔓茵科技的再生建材含杂率低，生产时进行了颗粒大小、生产成份种类的精准细分，在抗水、抗压等方面与普通建材一致，能够达到工程建设所用的标准。此外，再生建材的原材料符合绿色回收标准，能够节约部分建筑公司的工程成本。' where prod_t_id = 2;
update product_types set description = '再生骨料表面粗糙且母体混凝土块在经过本公司的设备进行破碎处理后，表面砂浆内部存在大量微裂纹，从而使其吸水率远高于天然骨料，且本公司对再生骨料进行了精准分类，根据不同的颗粒大小，继续针对性生产，提高了再生骨料的利用率，也间接提高了再生建材的整体性能。' where prod_t_id = 3;
update product_types set description = '蔓茵回收将不符合建筑工程施工标准的再生原材料进行多重利用，将其生产为再生工艺品，可放在公共场所或店面装饰，起到美化环境、提高资源利用率的作用。' where prod_t_id = 4;

update product_types set sub_text = '用于装配式建筑的一体化建筑设计和施工，配件可多次使用，减少垃圾产生' where prod_t_id = 1;
update product_types set sub_text = '蔓茵科技的再生材料含杂率低，原材料符合绿色回收标准，可节约工程成本' where prod_t_id = 2;
update product_types set sub_text = '根据颗粒大小对再生骨料进行了精准分类，继续针对性生产，提高了利用率，间接提高了再生建材的整体性能。' where prod_t_id = 3;
update product_types set sub_text = '将不符合标准的再生原材料进行多重利用，将其生产为再生工艺品' where prod_t_id = 4;

drop table product;

create table product (
    prod_id char(40) primary key,
    prod_name varchar(30) not null,
    prod_desc text not null,
    prod_type_id tinyint not null,
    prod_avatar text not null,
    prod_digital text not null,
    prod_price int(64) not null,
    prod_unit char(5) not null,
    foreign key (prod_type_id) references product_types(prod_t_id)
);

drop table paid_order;

create table paid_order (
    order_id bigint primary key,
    user_id  char(50),
    prod_id  char(40),
    prod_name  varchar(30) not null,
    prod_avatar text not null,
    total_price int(64) not null,
    buy_num int not null,
    addressee varchar(30),
    address  varchar(255),
    addressee_phone  varchar(20),
    create_time  timestamp not null,
    status  int not null,
    foreign key (user_id) references user(uuid),
    foreign key (prod_id) references product(prod_id)
);

insert into address values (
    _uuid,
    _id,
    _address,
    (
        case when
                     (
                         select a.num from
                             (select COUNT(1) as num from address where uuid = _uuid) as a
                     ) > 0
                 then 0
             else 1
            end
        ),
    (
        case when
                     (
                         select a.num from
                             (select COUNT(1) as num from address where uuid = _uuid) as a
                     ) <= 10
                 then 1
             else 0
            end
        )
);



-- insert into
-- c values
--       ('jack', case when (
--           select a.num from
--                            (select COUNT(1) as num from c group by name) as a
--           ) > 5
--           then 10
--           else -1
--           end
--       );
