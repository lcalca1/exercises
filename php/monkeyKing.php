 <?php
/************************************************************************
> File Name: yue.php
> Author: lca
> Created Time: Fri 04 Jan 2019 08:37:35 PM CST
> Desc：猴子选王，约瑟夫环
************************************************************************/

function _MonkeyKing($monkeys, $unlucky_n)
{
    $len = count( $monkeys );

    if ($len == 1) {
        return array( [$monkeys[0]], -1);
    }

    $unlucky_key = ($unlucky_n % $len) - 1;
    array_splice($monkeys, $unlucky_key, 1);
    return array($monkeys , $unlucky_key );
}

function MonkeyKing($monkeys, $unlucky_n )
{
    if (!is_array( $monkeys )) {
        return -1;
    }
    if (empty($monkeys)){
        return -1;
    }

    do {
        list($monkeys,) = _MonkeyKing($monkeys, $unlucky_n);
    } while( count($monkeys) > 1);

    return $monkeys[0];
}

$monkeys = [1,2,3,4,5,6,7,8,9,10];
$unlucky_n = 11;
print_r( MonkeyKing($monkeys, $unlucky_n) );
