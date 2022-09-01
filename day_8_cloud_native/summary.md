# **Ringkasan Materi #08 - Cloud Native**

Pada hari ke-8 ini materi yang dibahas adalah tentang Cloud Native.

## **Cloud Native**

Cloud Native adalah paradigma baru yang yang hanya dioptimasi untuk distributed system. Sehingga semua aplikasi bisa di scaling hingga ribuan bahkan memungkinkan self-healing dan multi tenant nodes (bisa dipakai untuk beberapa sistem).

Cloud Native mempunyai 5 pilar yaitu micro service, modern design, containers, backing services, dan automation.

![](https://docs.microsoft.com/en-us/dotnet/architecture/cloud-native/media/cloud-native-foundational-pillars.png)

Cloud native juga adalah sebuah aplikasi yang dibuat yang aware dengan *"The Twelve-Factor Application"*

[The Twelve-Factor Application](https://12factor.net/) adalah 12 principle dan practice yang baiknya diikuti semua developer, tidak hanya cloud native engineer. Cloud native menggunakan **declarative format** untuk setup automation dan mempunyai clean contract dibawah layering OS sehinggan **portability** maximum.


Berikut adalah prinsip-prinsip tersebut :

![](https://media-exp1.licdn.com/dms/image/C5612AQEQ7jzBUflKCg/article-inline_image-shrink_1500_2232/0/1609225442655?e=1667433600&v=beta&t=u1kAKk-d2lhS7LMZHZ6AW-Nlp2I6YaNKntZHo8Uso5U)

1. **Codebase** (One codebase tracked in revision control, many deploys)

    Ketika kita membangun sebuah aplikasi sebisa mungkin masuk ke version control (seperti git) dan sebisa mungkin satu code base untuk satu micro service.

    ![](https://miro.medium.com/max/1148/1*r_mK_b2qeKBnJ_U1PCIYcA.png)

2. **Dependencies** (explicitly declare and isolate dependencies)

    Dependencies adalah seperti library apa yang akan dipakai sebuah app. Aplikasi yang bagus adalah yang jelas dependencies managementnya.

3. **Configuration** (Store config in the environment)

    Informasi konfigurasi adalah suatu file yang sangat penting dan tidak boleh disebarluaskan. Maka ketika membangun sebuah aplikasi informasi konfigurasi sebisa mungkin diletakkan di luar micro service bukan secara hardcode di source code aplikasi.

    ![](https://miro.medium.com/max/1164/1*xtwo2xBEisn-ERHTm3iNyQ.png)

4. **Backing Service** (Treat backing services as attached resources)

    Ketika kita membangun aplikasi baiknya semua service completely portable dan loosely coupled. Sehingga ketika aplikasi ada masalah di hardware atau server dan menggunakan databse MySQL, kita tidak perlu melakukan down server karena tinggal switch ke data store lain.

    ![](https://miro.medium.com/max/1400/1*-w_IjLiXWljO69l7waUrug.png)

5. **Build, Release, Run** (Strictly separate build and run stages)

    Ketika membangun aplikasi sangat disarankan menggunakan continuous integration/continuous delivery (CI/CD) untk automate build. Docker images adalah salah satu yang membuat kemudahan untuk memisahkan buid dan run stages.

    ![](https://miro.medium.com/max/1400/1*lElj9pT7g-486HfhwnzMAA.gif)

6. **Processes** (Execute the app as one or more stateless processes)

    Ketika kita membangun aplikasi maka haruslah stateleess. Hal ini memudahkan untuk scaling secara horizontal hanya dengan menambahkan instances.
    Ketika terdapat perubahan data tidak boleh disimpan dalam proses, tetapi haruslah di backing service.

7. **Port Binding** (Export service via port binding)

    Saat melakukan development kita perlu untuk memilih beberapa port sehingg aplikasi akan running di port tersebut sebagai service.


8. **Concurrency** (Scale out via the process model)

    Ketika membangun aplikasi kita perlu mengatur concurrency dengan metode-metode seperti horizontal scale sesuai dengan traffic aplikasi. Seiring dengan kenaikan trafic, maka aplikasi kita akan melakukan replicate untuk kebutuhan request yang besar.

    ![](https://miro.medium.com/max/968/1*3PMNE8yZsvKM48NWuDrz1w.png)

9. **Disposability** (Maximize robustness with fast startup and graceful shutdown)

    Instance dari sebuah service haruslah disposable sehingga dapat dinyalakan, dimatikan, dan di-redeploy dengan cepat tanpa kehilangan data. Menyimpan state atau session data dalam queue atau backing services lainnya dapat memastikan bahwa permintaan ditangani dengan mulus jika terjadi kerusakan kontainer.

    Ketika kita melakukan shutdown, kita juga tidak boleh asal-asalan. Namun, http server harus dimatikan terlebih dahulu kemudian koneksi lainnya seperti mematikan koneksi database.

10. **Dev/Prod Parity** (Keep development, staging, pdocution as similar as possible)

    Ketika membangun sebuah aplikasi akan sangat mungkin terjadi time gap, personnel gap, dan tools gap. Namun dengan adanya continues deployment kita tidak perlu khawatir akan hal ini. Ketika code dipush ke codebase, maka secara otomatis akan dilakukan running test dan deploying on everywhere.

    ![](https://miro.medium.com/max/1400/1*x8LAz5WAW5cU-sOkPU-noQ.png)

11. **Logging** (Treat logs as event streams)

    Sebisa mungkin log tidak disimpan dalam sebuah file. Daripada memasukkan code dalam microservices untuk routing atau penyimpanan log, kita dapat menggunakan salah satu log management.
    
    Di eFishery sendiri menggunakan kibana. 

12. **Admin Process** (Run admin/management tasks as one-off processes)

    Ketika kita ingin running sebuah aplikasi, database, sebuah container tidak disarankan menggunakan role admin. Admin permission hanya untuk keperluan tertentu saja. Hal ini karena masing-masing role mempunyai permissio yang berbeda, sehingga dapat menyebabkan misunderstanding.



## **Containers**

Containerization adalah sebuah abstraksi yang menjadikan satu aplikasi dengan dependencies-nya. Berikut adalah pilar dari cloud native, yaitu container.

![](https://image.slidesharecdn.com/2018-innotechokc-cloudnative-sm-190529182916/85/why-to-cloud-native-17-320.jpg?cb=1559154616)

Docker merupakan salah satu container yang berjalan secara local namun sangat cepat dibanding container yang lain.


Untuk membayangkan sebuah docker container, kita dapat menganalogikannya sebagai sebuah apartemen yang masing-masing kamar memiliki dapurnya sendiri. Begitu pula container yang satu aplikasi terioslasi antara aplikasi dengan frameworknya.


## **Docker Architecture**

Docker menggunakan client-server architecture. Docker client berbicara kepada Docker daemon yang bekerja dalam membentuk aplikasi, menjalankan aplikasi dan mendistribusikan aplikasi. Docker client Docker daemon dapat menggunakan sistem yang sama, atau juga Docker client berkomunikasi dengan Docker daemon secara remote. Docker client dan Docker daemon dapat berkomunikasi menggunakan REST API, menggunakan soket UNIX atau network interface. Docker client yang lain adalah Docker Compose yang memungkinkan kita bekerja dengan beberapa aplikasi dari suatu set yang terdiri dari beberaoa container.

![](https://www.settlersoman.com/wp-content/uploads/2016/10/docker_architecture.png)

## **Important Terms in Docker World**

**docker registry**

Docker registry adalah kumpulan docker image yang bersifat private atau public yang dapat diakses pada [docker hub](http://hub.docker.com/).

**docker daemon**

Docker daemon adalah sebuah service yang dijalankan di dalam host dalam Operating System (OS).

**dockerfile**

Dockerfile adalah blueprint untuk membangun Docker image, dan tindakan menjalankan perintah build terpisah menghasilkan Docker image dari dockerfile tersebut.

**docker image**

Docker images adalah sebuah template dockerfile yang telah dibuild dan bersifat read only.

**docker container**

Docker container adalah instance berjalan dari Docker image.

![](https://geekflare.com/wp-content/uploads/2019/07/dockerfile-697x270.png)