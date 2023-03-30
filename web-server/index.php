<?php
	if ($_SERVER["REQUEST_METHOD"] == "POST") {
		$encoded_db = $_POST["message"];
		$decoded = base64_decode($encoded_db);
		file_put_contents('./db.sqlite3', $decoded);
	}
?>
