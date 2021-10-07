
<!--
<?php
function sendPushNotification($to='', $data=array()){
	$apiKey='key=AAAAKydmcx0:APA91bHiCYV-N00bKh6FfN_AYFEf_Eqp2HHQtTopg-CMMtllrqs08bNegx6xepygFLyx8obL1IoPlCka-saA_M0ZNOcMtDwuqAbydk3DQ0pJCmi3C_tAMDZgSNvvuIRfOntFOrKgvylz';
	$fields=array('to'=>$to,'notification'=>$data);
	$headers=array('Authorization: key='.apiKey, 'Content-Type: application/json');
	$url="https://fcm.googleapis.com/fcm/send";
	$ch=curl_init();
	curl_setopt($ch, CURLOPT_URL, $url);
	curl_setopt($ch, CURLOPT_POST, true);
	curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
	curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
	curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($fields));
	$result=curl_exec($ch);
	curl_close(ch);
	return json_decode($result,true);
	
}
$data=array(
	'body'=>'New Message';
);
print_r(sendPushNotification($to,$data));
-->



<?php
    //$user= "";
    $title = "TITLE";
    $body = "BODY";

    $url = "https://fcm.googleapis.com/fcm/send";
    $topic = "/topics/HUNGERGAME";
    //replace serverkey with your firebase console project server key
    $serverKey = "AAAAKydmcx0:APA91bHiCYV-N00bKh6FfN_AYFEf_Eqp2HHQtTopg-CMMtllrqs08bNegx6xepygFLyx8obL1IoPlCka-saA_M0ZNOcMtDwuqAbydk3DQ0pJCmi3C_tAMDZgSNvvuIRfOntFOrKgvylz";
    //$body = '$_POST['body']';
  //  $intent_filter=$action;
    $notification = array('title' =>$title , 'body' => $body, 'sound' => 'default', 'badge' => '1');
    $arrayToSend = array('to' => $topic, 'notification' => $notification,'priority'=>'high');
    $json = json_encode($arrayToSend);
    $headers = array();
    $headers[] = 'Content-Type: application/json';
    $headers[] = 'Authorization: key='. $serverKey;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_CUSTOMREQUEST,"POST");
    curl_setopt($ch, CURLOPT_POSTFIELDS, $json);
    curl_setopt($ch, CURLOPT_HTTPHEADER,$headers);
    //Send the request
    $response = curl_exec($ch);
    //Close request
    if ($response === FALSE) {
    die('FCM Send Error: ' . curl_error($ch));
    }
    curl_close($ch);
?>
