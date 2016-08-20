#!/bin/bash

# author: yingx
# date: 2016-08-18

function usage() {
    echo "USAGE: $0 create|remove modulename templatename"
    exit 1
}

function capitalize() {
    split=( $1 )
    echo "${split[*]^}"
}

function substr() {
    local str1=$1
    local str2=$2

    if [[ ${str1/${str2}//} == $str1 ]]; then
        return 0
    else
        return 1
    fi
}

function checkParams() {
    case "$1" in
        "create")
            if [ ! -d $5/$2 ]; then
                echo "The template:$2 does not exist!"
                echo ""
                usage
            fi
            if [ -f $4/${3}Controller.go ]; then
                echo "The module:$3 has already existed, please try another name!"
                echo ""
                usage
            fi
            ;;
        "remove")
            ;;
        *)
            echo "The operation:$1 is not supported!"
            echo ""
            usage
            ;;
    esac
}

function createModule() {
    echo "USAGE: $0 create|remove templatename modulename"
}

function removeModule() {
    echo "USAGE: $0 create|remove templatename modulename"
}

#retrieve params
case "$#" in
    "2")
        operation=$1
        modulename=$2
        templatename=basic
        ;;
    "3")
        operation=$1
        modulename=$2
        templatename=$3
        ;;
    *)
        echo "Parameter count is not right!"
        echo ""
        usage
        ;;
esac

basedir=`pwd`
controllerpath=$basedir/server/controller
modelpath=$basedir/server/model
datapath=$basedir/server/data

templatespath=$basedir/templates
templatescontrollerpath=$templatespath/$templatename/server/controller
templatesmodelpath=$templatespath/$templatename/server/model
templatesdatapath=$templatespath/$templatename/server/data

#check params
checkParams $operation $templatename $modulename $controllerpath $templatespath

case "$operation" in
    "create")
        echo "Creating a new module $modulename by template $templatename"
        templatenameCap=`capitalize $templatename`
        modulenameCap=`capitalize $modulename`

        moduleController=$controllerpath/${modulename}Controller.go
        echo "Cteate $moduleController ... "
        cp $templatescontrollerpath/${templatename}Controller.go $moduleController
        sed -i "s/$templatename/$modulename/g" $moduleController
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleController

        moduleSrvController=$controllerpath/${modulename}SrvController.go
        echo "Cteate $moduleSrvController ... "
        cp $templatescontrollerpath/${templatename}SrvController.go $moduleSrvController
        sed -i "s/$templatename/$modulename/g" $moduleController
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleController

        index=0
        sqlflag=0
        createsqlitems=""
        selectsqlitems=""
        insertsqlcolumn=""
        insertsqlvalue=""
        updatesqlitems=""
        jsonitems="    "
        templateSrvData=$templatesdatapath/${templatename}SrvData.txt
        while IFS='|' read -ra sqlcolumn; do
            if [ $index -ne 0 ]; then
                createsqlitems=${createsqlitems},
                selectsqlitems=${selectsqlitems},
                jsonitems="${jsonitems}\n    "
            fi
            for info in "${sqlcolumn[@]}"; do
                createsqlitems="${createsqlitems} ${info}"
            done
            if [ $sqlflag -ne 0 ]; then
                insertsqlcolumn=${insertsqlcolumn},
                insertsqlvalue=${insertsqlvalue},
                updatesqlitems=${updatesqlitems},
            fi
            selectsqlitems="${selectsqlitems} ${sqlcolumn[0]}"
            if [ ${sqlcolumn[0]} != "id" ]; then
                insertsqlcolumn="${insertsqlcolumn} ${sqlcolumn[0]}"
                insertsqlvalue="${insertsqlvalue} ?"
                updatesqlitems="${updatesqlitems} ${sqlcolumn[0]}=?"
                sqlflag=1
            fi
            columnNameCap=`capitalize ${sqlcolumn[0]}`
            substr ${sqlcolumn[1]} "char"
            if [ $? -eq 1 ]; then
                jsonitems="${jsonitems} ${columnNameCap}  string \`json:\"${sqlcolumn[0]}\"\`"
            else
                jsonitems="${jsonitems} ${columnNameCap} int64 \`json:\"${sqlcolumn[0]}\"\`"
            fi
            index=$(( $index + 1 ))
        done < $templateSrvData
        
        createsql="create table ${modulename} ($createsqlitems)"
        moduleSrvData=$datapath/${modulename}SrvData.sql
        echo "Cteate $moduleSrvData ... "
        echo $createsql > $moduleSrvData

        moduleSrvModel=$modelpath/${modulename}SrvModel.go
        echo "Cteate $moduleSrvModel ... "
        cp $templatesmodelpath/${templatename}SrvModel.go $moduleSrvModel
        sed -i "s/$templatename/$modulename/g" $moduleSrvModel
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleSrvModel
        sed -i "s/##Item##/${jsonitems}/g" $moduleSrvModel
        selectsql="select${selectsqlitems} from ${modulename}"
        selectonesql="select${selectsqlitems} from ${modulename} where id=?"
        insertsql="insert into ${modulename}(${insertsqlcolumn}) values(${insertsqlvalue})"
        updatesql="update ${modulename} set${updatesqlitems} where id=?"
        deletesql="delete from ${modulename} where id=?"
        sed -i "s/##query##/${selectsql}/g" $moduleSrvModel
        sed -i "s/##queryone##/${selectonesql}/g" $moduleSrvModel
        sed -i "s/##insert##/${insertsql}/g" $moduleSrvModel
        sed -i "s/##update##/${updatesql}/g" $moduleSrvModel
        sed -i "s/##delete##/${deletesql}/g" $moduleSrvModel
        ;;
    "remove")
        echo "Remove the existed module $modulename"
        rm $controllerpath/${modulename}Controller.go
        rm $controllerpath/${modulename}SrvController.go
        rm $modelpath/${modulename}SrvModel.go
        rm $datapath/${modulename}SrvData.sql
        ;;
esac



