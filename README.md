# Cookie Monster

Welcome to "Cookie Monster"! This software is designed to be installed by a victim, at the discretion of the attacker. Once installed, "Cookie Monster" will locate the victim's sqlite3 database containing their session cookies and encrypt it. The encrypted file will then be sent over an HTTP connection back to the attackers web server. It is important to note that this software is intended for educational purposes only and should not be used for any illegal or malicious activities. Please use this software responsibly and ensure that you have the necessary permissions before using it on any system. 

## Background

By default, the "Cookies" sqlite3 file used by Google Chrome to store session cookies is not encrypted. This means that if an attacker gains access to this file, they could potentially read the data stored in the users cookies, including sensitive information such as login credentials and session tokens. This repository demonstrates just that.


## Here's how it works

The Target downloads the CookieMonster binary executable, Cookie-Monster finds the session cookies, encrypts the contents with AES-256, then base64 encodes the encrypted file. Cookie-Monster then relays the base64 string back to the attackers web server.


## Chrome cookie locations:

<h3>macOS</h3>
`/Users/[user]/Library/Application Support/Google/Chrome/Default/Cookies`

<h3>Linux</h3>
`/home/[user]/.config/google-chrome/Default/Cookies`

<h3>Windows</h3>
`C:\Users[user]\AppData\Local\Google\Chrome\User Data\Default\Cookies`




<h2>Disclaimer:</h2>

The following GitHub repository contains a methodology for attacking applications. The methodology is provided for educational and research purposes only, and should not be used for any illegal or malicious activities.

By accessing or using the methodology, you acknowledge and accept that you do so at your own risk. The authors and contributors of this repository are not responsible for any damages, losses, or legal consequences that may result from the use of this methodology.

Furthermore, the authors and contributors of this repository do not endorse or condone any illegal or unethical activities that may be facilitated by the use of this methodology. It is your responsibility to ensure that your use of the methodology is in compliance with all applicable laws and regulations.

Please note that the use of this methodology on systems or networks without proper authorization may be considered a criminal offense. You are solely responsible for obtaining any necessary permissions or authorizations before using the methodology.

The authors and contributors of this repository reserve the right to remove, modify, or discontinue the methodology at any time without notice.

By accessing or using the methodology, you agree to indemnify, defend, and hold harmless the authors and contributors of this repository from any claims, damages, losses, or expenses arising out of your use of the methodology or your violation of this disclaimer.

Use of this methodology is strictly at your own risk, and the authors and contributors of this repository are not responsible for any consequences that may result from its use.
