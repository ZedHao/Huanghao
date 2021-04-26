
1.bst的创建和递归遍历
```
<?php

class Node {
    public $key;
    public $left;
    public $right;

    public function __construct($key) {
        $this->key    = $key;
        $this->left   = null;
        $this->right  = null;
    }
}

class Bst {
    public $root;

    /**
     * 初始化树结构
     *
     * @param $arr 初始化树结构的数组
     *
     * @return null
     */
    public function init($arr) {
        $this->root = new Node($arr[0]);
        for($i = 1; $i < count($arr); $i++) {
            $ret = $this->Insert($arr[$i]);
        }
        return $ret;
    }

    /**
     * 将$key插入树中
     *
     * @param $key 待插入树的数字
     *
     * @return null
     */
    function Insert($key) {
        if(!is_null($this->search($key))){
            throw new  Exception();
        }
        $root = $this->root;
        $innode = new Node($key);
        $current = $root;
        $preNode = null;

        //先找到适合key插入的节点
        while($current !=null){
            $preNode = $current;
            if($current->key>$innode->key){
                $current = $current->left;
            }else{
                $current = $current->right;
            }
        }
        //将key 插入到找出的节点左右
        if($preNode == null){
            $this->root = $innode;
        }else{
            if($preNode->key > $innode->key){
                $preNode->left = $innode;
            }else{
                $preNode->right = $innode;
            }
        }
        return $root;
    }
    /**
     * 前序遍历二叉树
     * @param $node
     * User: huanghao@zuoyebang.com
     * Date: 2019/3/7
     * Time: 11:59
     */
    function preTransTree($node,$string){

        if(!is_null($node->left)){
            $string = $this->preTransTree($node->left,$string);
        }
        if(!is_null($node->key)){
            $string .=$node->key."->";
        }
        if(!is_null($node->right)){
            $string = $this->preTransTree($node->right,$string);
        }

        return $string;
    }
    /**
     * 中序遍历二叉树
     * @param $node
     * User: huanghao@zuoyebang.com
     * Date: 2019/3/7
     * Time: 11:59
     */
    function midTransTree($node,$string){
        if(!is_null($node->key)){
            $string .=$node->key."->";
        }
        if(!is_null($node->left)){
            $string = $this->midTransTree($node->left,$string);
        }
        if(!is_null($node->right)){
            $string = $this->midTransTree($node->right,$string);
        }

        return $string;
    }
    /**
     * 后序遍历二叉树
     * @param $node
     * User: huanghao@zuoyebang.com
     * Date: 2019/3/7
     * Time: 11:59
     */
    function latTransTree($node,$string){

        if(!is_null($node->left)){
            $string = $this->latTransTree($node->left,$string);
        }
        if(!is_null($node->right)){
            $string = $this->latTransTree($node->right,$string);
        }
        if(!is_null($node->key)){
            $string .=$node->key."->";
        }

        return $string;
    }


    /**
     * 查找树中是否存在$key对应的节点
     *
     * @param $key 待搜索数字
     *
     * @return $key对应的节点
     */
    function search($key) {
        $current = $this->root;
        while($current != null) {
            if($current->key == $key) {
                return $current;
            } elseif($current->key > $key) {
                $current = $current->left;
            } else {
                $current = $current->right;
            }
        }

        return $current;
    }

}
$arr = array(5, 3, 8, 2, 4, 7, 9, 1, 6);
$obj = new Bst();
$ret = $obj->init($arr);
$string = "前序".$obj->preTransTree($obj->root,"");
var_dump($string);
$string = "中序".$obj->midTransTree($obj->root,"");
var_dump($string);
$string = "后序".$obj->latTransTree($obj->root,"");
var_dump($string);
//

```
```
