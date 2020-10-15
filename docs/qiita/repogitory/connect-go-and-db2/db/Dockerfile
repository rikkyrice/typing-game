FROM ibmcom/db2:11.5.4.0
RUN yum -y update && yum clean all
# 言語設定
RUN sed -i -e '/override_install_langs/s/$/,ja_JP.utf8/g' /etc/yum.conf
## デフォルトで
RUN curl http://mirror.centos.org/centos/7/os/x86_64/Packages/glibc-common-2.17-307.el7.1.x86_64.rpm > /tmp/glibc-common.rpm
RUN rpm -Uvh --force /tmp/glibc-common.rpm
RUN yum -y install ibus-kkc vlgothic-fonts vlgothic-p-fonts
RUN localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
RUN echo 'LANG="ja_JP.UTF-8"' >  /etc/locale.conf
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
# 時刻設定
RUN echo 'ZONE="Asia/Tokyo"' > /etc/sysconfig/clock
RUN rm -f /etc/localtime
RUN ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN mkdir /var/custom
RUN mkdir /var/custom/sql
RUN mkdir /var/custom/data
COPY ./sql/*.sql /var/custom/sql/
COPY ./data/*.csv /var/custom/data/
COPY createschema.sh /var/custom/
RUN chmod 444 /var/custom/sql/*.sql
RUN chmod 444 /var/custom/data/*.csv
RUN chmod a+x /var/custom/createschema.sh