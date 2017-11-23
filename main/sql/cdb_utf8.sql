drop table picc_verification;
drop table picc_role_privilege;
drop table picc_user_role;
drop table picc_loss;

drop table picc_subject;
drop table picc_insurance;
drop table picc_role;
drop table picc_entity;
drop table picc_user;
drop table picc_insured;


drop table picc_subject_attr;
/*==============================================================*/
/* Table: picc_entity                                           */
/*==============================================================*/
create table picc_entity (
   id                   SERIAL               not null,
   name                 VARCHAR(2)           not null unique,
   alias                VARCHAR(128)         not null,
   constraint PK_PICC_ENTITY primary key (id)
);

comment on column picc_entity.name is
'应该添加unique约束，且为两个符';

/*==============================================================*/
/* Table: picc_insurance                                        */
/*==============================================================*/
create table picc_insurance (
   id                   varchar(128)         not null,
   insured_pay          varchar(256) not null,
   insured_benefit      varchar(256) not null,
   relation             relation_type           not null,
   type                 insurance_type           not null,
   households           INT4                 null,
   insurance_amount     MONEY                not null,
   insurance_fee        MONEY                not null,
   insurance_fee_self   MONEY                null,
   dispute              VARCHAR(128)         null,
   time_stamp           TIMESTAMP            not null,
   undertake            VARCHAR(16)          not null,
   constraint PK_PICC_INSURANCE primary key (id)
);

comment on table picc_insurance is
'作为表单表管理所有的的表单';

comment on column picc_insurance.relation is
'[enum] 投保人与被投保人关系  本人、村民和村委';

comment on column picc_insurance.type is
'[enum] 投保方式 个人投保，团体投保';

comment on column picc_insurance.households is
'户数';

comment on column picc_insurance.insurance_amount is
'总保险金额';

comment on column picc_insurance.insurance_fee is
'保险费';

comment on column picc_insurance.undertake is
'承保人';

/*==============================================================*/
/* Table: picc_insured                                          */
/*==============================================================*/
create table picc_insured (
   name                 varchar(256)         not null,
   code                 varchar(128)         not null,
   tel                  varchar(32)          not null,
   prov                 varchar(64)          null,
   city                 varchar(64)          not null,
   county               varchar(64)          not null,
   town                 varchar(64)          not null,
   village              varchar(64)          not null,
   zip_code             varchar(32)          null,
   post_code            varchar(32)          null,
   constraint PK_PICC_INSURED primary key (name)
);

comment on table picc_insured is
'投保人和被投保人都用一个表来表示，主键为组织机构代码/身份证号';

comment on column picc_insured.code is
'组织机构代码或省份证号';

/*==============================================================*/
/* Table: picc_loss                                             */
/*==============================================================*/
create table picc_loss (
   id                   INT4                 not null,
   staff                VARCHAR(16)          null,
   subject_id           INT4                 null,
   time_stamp           DATE                 not null,
   loss_level           INT4                 not null,
   pic1                 CHAR(254)            null,
   pic2                 CHAR(254)            null,
   pic3                 CHAR(254)            null,
   pic4                 CHAR(254)            null,
   pic5                 CHAR(254)            null,
   constraint PK_PICC_LOSS primary key (id)
);

comment on table picc_loss is
'定损相关信息放在此处 损失等级可以用enum表示';

/*==============================================================*/
/* Table: picc_role                                             */
/*==============================================================*/
create table picc_role (
   id                   SERIAL               not null,
   name                 VARCHAR(128)         null,
   constraint PK_PICC_ROLE primary key (id)
);

/*==============================================================*/
/* Table: picc_role_privilege                                   */
/*==============================================================*/
create table picc_role_privilege (
   id                   SERIAL               not null,
   role_id              INT4                 null,
   entity_id            INT4                 null,
   can_add              INT4                 null,
   can_read             INT4                 null,
   can_update           INT4                 null,
   can_del              INT4                 null,
   can_assign           INT4                 null,
   constraint PK_PICC_ROLE_PRIVILEGE primary key (id)
);

/*==============================================================*/
/* Table: picc_subject                                          */
/*==============================================================*/
create table picc_subject (
   id                   SERIAL               not null,
   insurance_id         VARCHAR(128)         null,
   subject_name         VARCHAR(128)         not null,
   unit_insurance_fee   MONEY                not null,
   insurance_amount     INT4                 not null,
   insurance_deductible MONEY                null,
   insurance_deductible_percent FLOAT8               null,
   insurance_fee_total  MONEY                not null,
   insuracne_fee_central INT4                 null,
   insurance_fee_prov   INT4                 null,
   insurance_fee_city   INT4                 null,
   insurance_fee_county INT4                 null,
   insurance_fee_self   INT4                 null,
   insurance_period_start DATE                 not null,
   insurance_period_end DATE                 not null,
   status               insurance_status_type   default '正常',
   subject              json           not null,
   area                 polygon           null,
   constraint PK_PICC_SUBJECT primary key (id)
);

comment on table picc_subject is
'该表为保险标的表，其中subject字段用json设置，存储不同标的需要考察的要素，要素的key由picc_subject_mgr表确定，还有一个字段存储地块的geo信息。';

comment on column picc_subject.subject_name is
'标的名称';

comment on column picc_subject.unit_insurance_fee is
'单位保险金额';

comment on column picc_subject.insurance_amount is
'保险数量';

comment on column picc_subject.insurance_deductible is
'免赔额';

comment on column picc_subject.insurance_deductible_percent is
'免赔率';

comment on column picc_subject.insurance_fee_total is
'总保费';

comment on column picc_subject.insuracne_fee_central is
'中央承担的保费';

comment on column picc_subject.insurance_fee_prov is
'省承担的保费';

comment on column picc_subject.insurance_fee_city is
'市承担的保费';

comment on column picc_subject.insurance_fee_county is
'县承担的保费';

comment on column picc_subject.insurance_fee_self is
'自担部分';

comment on column picc_subject.status is
'保单状态，正常，审核中，已审核，出险中，出险完毕，过期';

/*==============================================================*/
/* Table: picc_subject_attr                                     */
/*==============================================================*/
create table picc_subject_attr (
   id                   SERIAL               not null,
   subject_name         VARCHAR(128)         not null,
   subject_col          VARCHAR(128)         not null,
   constraint PK_PICC_SUBJECT_ATTR primary key (id)
);

comment on table picc_subject_attr is
'保险标的表的管理表，保险标的字段存储在该表中，由于每种保险产品的标的具有不同的属性，所以定义该表，该表的作用是确定保险标的属性。在保险标的表中，存在一个可变字段，用json来表示，其中json的key就是该表的subject_col确定 
不再定义规则表，也就是说单价直接输入，json中的内容，仅作注释参考用';

/*==============================================================*/
/* Table: picc_user                                             */
/*==============================================================*/
create table picc_user (
   id                   VARCHAR(32)          not null,
   tel                  VARCHAR(32)          not null unique,
   name                 VARCHAR(32)          not null,
   passwd               VARCHAR(128)         not null,
   prov                 VARCHAR(64)          not null,
   city                 VARCHAR(64)          not null,
   county               VARCHAR(64)          not null,
   post_code            VARCHAR(32)          not null,
   create_time          TIMESTAMP            default now(),
   constraint PK_PICC_USER primary key (id)
);
   /*constraint AK_KEY_UNIQUE_PICC_USE unique ()
   constraint AK_KEY_3_PICC_USE unique ()
   constraint AK_KEY_4_PICC_USE unique ()
   constraint AK_KEY_5_PICC_USE unique ()
*/

comment on table picc_user is
'管理picc工作人员';

comment on column picc_user.tel is
'应该添加Unique约束';

comment on column picc_user.create_time is
'默认值当前时间';

/*==============================================================*/
/* Table: picc_user_role                                        */
/*==============================================================*/
create table picc_user_role (
   id                   SERIAL               not null,
   role_id              INT4                 null,
   user_id              VARCHAR(32)          null,
   constraint PK_PICC_USER_ROLE primary key (id)
);

/*==============================================================*/
/* Table: picc_verification                                     */
/*==============================================================*/
create table picc_verification (
   id                   SERIAL               not null,
   staff                VARCHAR(16)          not null,
   insurance            VARCHAR(128)         not null,
   result               INT4                 not null,
   result_desc          VARCHAR(64)          not null,
   time_stamp           DATE                 not null,
   constraint PK_PICC_VERIFICATION primary key (id)
);

comment on table picc_verification is
'核保信息表';

comment on column picc_verification.result is
'0 表示核保通过';

alter table picc_insurance
   add constraint FK_PICC_INS_REFERENCE_PICC_INSP foreign key (insured_pay)
      references picc_insured (name)
      on delete restrict on update restrict;

alter table picc_insurance
   add constraint FK_PICC_INS_REFERENCE_PICC_INSB foreign key (insured_benefit)
      references picc_insured (name)
      on delete restrict on update restrict;

alter table picc_insurance
   add constraint FK_PICC_INS_REFERENCE_PICC_USE foreign key (undertake)
      references picc_user (id)
      on delete restrict on update restrict;

alter table picc_loss
   add constraint FK_PICC_LOS_REFERENCE_PICC_USE foreign key (staff)
      references picc_user (id)
      on delete restrict on update restrict;

alter table picc_loss
   add constraint FK_PICC_LOS_REFERENCE_PICC_SUB foreign key (subject_id)
      references picc_subject (id)
      on delete restrict on update restrict;

alter table picc_role_privilege
   add constraint FK_PICC_ROL_REFERENCE_PICC_ROL foreign key (role_id)
      references picc_role (id)
      on delete restrict on update restrict;

alter table picc_role_privilege
   add constraint FK_PICC_ROL_REFERENCE_PICC_ENT foreign key (entity_id)
      references picc_entity (id)
      on delete restrict on update restrict;

alter table picc_subject
   add constraint FK_PICC_SUB_REFERENCE_PICC_INS foreign key (insurance_id)
      references picc_insurance (id)
      on delete restrict on update restrict;

alter table picc_user_role
   add constraint FK_PICC_USE_REFERENCE_PICC_ROL foreign key (role_id)
      references picc_role (id)
      on delete restrict on update restrict;

alter table picc_user_role
   add constraint FK_PICC_USE_REFERENCE_PICC_USE foreign key (user_id)
      references picc_user (id)
      on delete restrict on update restrict;

alter table picc_verification
   add constraint FK_PICC_VER_REFERENCE_PICC_USE foreign key (staff)
      references picc_user (id)
      on delete restrict on update restrict;

alter table picc_verification
   add constraint FK_PICC_VER_REFERENCE_PICC_INS foreign key (insurance)
      references picc_insurance (id)
      on delete restrict on update restrict;

