<?php
    include_once "db_tools.php";

    $host = "localhost";
    $user = "root";
    $pwd = "";
    $db = "nodeMCU";
    
    try{
        //建立資料庫
        $conn = create_conn($host, $user, $pwd, $db);
        //新增資料
        try{
            $sql = "SELECT * FROM temp_humi";
            $result = read_data($conn, $sql);
            $col_name = mysqli_fetch_field($result);
            $rows = mysqli_fetch_all($result, MYSQLI_ASSOC);

            $data = json_encode($rows);
            echo $data;
            /*array_map("myfunction",$rows);
            
            function array_json($data, $col){
                $output[$col] = jason_encode();
                return $output;
            }*/
            mysqli_free_result($result);

        }catch(Exception $err){
            echo "mesage: ".$err->getMessage();                
            die();
        }
    }catch(Exception $err){
        echo "mesage: ".$err->getMessage();                
        die();
    }
?>