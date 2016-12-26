:: 关闭终端回显
@echo off

set "IDCard_File=IDCard.xlsx"
set "IDCard_Sheet=IDCard "
set "TEMPLATE_File=入职人员信息表-模板.xlsx"
set "TEMPLATE_Sheet=入职信息表"
set "DeptInfo_File=部门编号.xlsx"
set "DeptInfo_Sheet=Sheet1"
set "Position_No=10000471"
set "Dimission_ID_File=Dimission_ID.xlsx"
set "Dimission_ID_Sheet=Sheet1"

employee_list.exe

pause

