FROM php:fpm

# bcmath bz2 calendar ctype curl dba dom enchant exif ffi fileinfo filter ftp gd gettext gmp 
#hash iconv imap intl json ldap mbstring mysqli oci8 odbc opcache pcntl pdo pdo_dblib 
#pdo_firebird pdo_mysql pdo_oci pdo_odbc pdo_pgsql pdo_sqlite pgsql phar posix pspell readline reflection 
#session shmop simplexml snmp soap sockets sodium spl standard sysvmsg sysvsem sysvshm tidy
# tokenizer xml xmlreader xmlrpc xmlwriter xsl zend_test zip
RUN /usr/local/bin/docker-php-ext-install pdo pdo_mysql

RUN mkdir /www/hello -p

ADD ./conf/www.conf /usr/local/etc/php-fpm.d/www.conf

ADD  ./conf/php.ini /usr/local/etc/php/php.ini

EXPOSE 9000/tcp

