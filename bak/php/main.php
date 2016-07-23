<?php

function add(&$c, $a, $b) {
	$c = $a + $b;
}

$a = 0;
$b = 0;
$c = 0;
$a = 2;
$b = 3;
add($c, $a, $b);
echo $c;
