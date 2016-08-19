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

        templateSrvData=$templatesdatapath/${templatename}SrvData.txt
        moduleSrvData=$datapath/${modulename}SrvData.sql
        echo "Cteate $moduleSrvData ... "
        #touch $moduleSrvData
        totalline=`awk -F"|" 'END {print(NR)}' $templateSrvData`
        sqlitems=`awk -F"|" '{if(NR != "'$totalline'") printf("%s %s %s,", $1, $2, $3); else printf("%s %s %s", $1, $2, $3);}' $templateSrvData`
        createsql="create table ${modulename} ($sqlitems)"
        echo $createsql > $moduleSrvData

        moduleSrvModel=$modelpath/${modulename}SrvModel.go
        echo "Cteate $moduleSrvModel ... "
        cp $templatesmodelpath/${templatename}SrvModel.go $moduleSrvModel
        sed -i "s/$templatename/$modulename/g" $moduleSrvModel
        sed -i "s/$templatenameCap/$modulenameCap/g" $moduleSrvModel
        items=`awk -F"|" '{if($2 ~ "char") printf("%s %s \\\\\`json:\"%s\"\\\\\`\\n", $1, "string", $1); else printf("%s %s \\\\\`json:\"%s\"\\\\\`\\n", $1, "int64", $1);}' $templateSrvData`
        ITEM="##Item##"
        echo $ITEM
        echo $items
        sed -i "s/${ITEM}/```${items}```/g" $moduleSrvModel
        ;;
    "remove")
        echo "Remove the existed module $modulename"
        rm $controllerpath/${modulename}Controller.go
        rm $controllerpath/${modulename}SrvController.go
        rm $modelpath/${modulename}SrvModel.go
        rm $datapath/${modulename}SrvData.sql
        ;;
esac



