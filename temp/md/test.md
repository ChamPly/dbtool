# XXXX系统

### 一、媒资信息

#### 1. 数据字典 [epg_sys_dictionary]
|字段名|类型|默认值|为空|约束|描述|
|---|---|---|---|---|---|
|id|number(20)|-|否|PK|编号|
|value|varchar(32)|-|否||枚举项的值|
|description|varchar(64)|-|否||枚举项的描述|
|belong_enum|varchar(32)|-|否||所属枚举|
|sort_id|number(2)|0|否||枚举内排序|
|group_id|varchar(32)|-|是||分组|
|status|number(2)|0|否||状态|
|remark|varchar(64)|-|是||备注|