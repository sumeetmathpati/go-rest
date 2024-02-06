grammar http;

http_message: request | response;

request: request_line CRLF headers CRLF body?;
response: status_line CRLF headers CRLF body?;

request_line: method SP request_target SP http_version;
status_line: http_version SP status_code SP reason_phrase;

method: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'HEAD' | 'PATCH' | 'OPTIONS' | 'TRACE';
request_target: '/' path | absolute_uri;
path: '/' segment;
segment: pchar+;
absolute_uri: scheme '://' host (':' port)? path ('?' query)?;
scheme: 'http' | 'https';
host: hostname | ipv4_address;
hostname: (ALPHA | DIGIT | '-' | '.')+;
ipv4_address: DIGIT '.' DIGIT '.' DIGIT '.' DIGIT;
port: DIGIT+;
query: (pchar | '/' | '?')+;

http_version: 'HTTP/' DIGIT '.' DIGIT;
status_code: DIGIT DIGIT DIGIT;
reason_phrase: TEXT;

headers: header*;
header: field_name ':' field_value;
field_name: token;
field_value: TEXT | token;

body: (TEXT | OCTET)*;

SP: ' ';
CRLF: '\r\n';

// Token Definitions
DIGIT: [0-9];
ALPHA: [a-zA-Z];
pchar: unreserved | escaped | ':' | '@' | '&' | '=' | '+' | '$' | ',' | '/' | '?' | ';' | '%';
unreserved: ALPHA | DIGIT | '-' | '.' | '_' | '~';
escaped: '%' HEXDIGIT HEXDIGIT;
HEXDIGIT: [0-9a-fA-F];
TEXT: ~[\x00-\x1F\x7F()<>@,;:\\\"/\[\]?={}\x20\t]+;
token: ~[\x00-\x1F\x7F()<>@,;:\\\"/\[\]?={}\x20\t]+;
OCTET: '\x00'..'\xFF';
