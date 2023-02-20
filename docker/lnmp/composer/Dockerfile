FROM composer

RUN mkdir /data

WORKDIR /data

RUN composer config -g repo.packagist composer https://packagist.phpcomposer.com

# RUN composer global require laravel/installer

ENTRYPOINT ["composer", "create-project", "--prefer-dist", "laravel/laravel", "laravel" ]