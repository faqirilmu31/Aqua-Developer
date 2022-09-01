# **Ringkasan Materi #09 - Cloud Native 2**

Pada hari ke-9 ini materi yang dibahas adalah Docker Orchestration dan CI/CD


## **Docker Orchestration**

**Container Orchestration** adalah sebuah automation tools untuk mengurangi sebagian besar upaya operasional yang diperlukan untuk menjalankan beban kerja dan layanan dalam container provisioning, deployment, scaling, networking, load balancing, dan sebagainya. Sedangkan **Docker Orchestration** adalah seperangkat praktik dan teknologi untuk mengelola Docker container dalam skala besar.

Ada beberapa tools untuk container orchestration seperti Kubernetes, Docker Swarm, Amazon Elastic Container Service (EKS), Google Cloud Run, Azure Container Instances (ACI), dan lain-lain.

Perbedaan antara docker dengan docker compose adalah :
- docker hanya menjalankan container satu per satu
- sedangkan docker compose memungkinkan semua kontainer berjalan bersama di satu waktu

**Manfaat menggunakan container orchestration :**

- simplified operations

  Container dibentuk dengan kompleksitas yang tinggi yang dapat dengan cepat lepas kendali. Oleh karena itu, simplified operations bermanfaat untuk mengelolanya.

- resilience

  Container orchestration tools dapat mengotomasi restart atau scale sebuah container atau cluster.

- adder security

  Container orchestration membantu mengurangi atau menghilangkan kemungkinan kesalahan manusia.


## **Docker Compose**

[Docker Compose](https://docs.docker.com/compose/) adalah alat untuk menjalankan satu atau beberapa container yang saling terkait dengan sebuah command.

![](https://i0.wp.com/i.ibb.co/CHzcf7s/docker-compose-logo.png?resize=810%2C398&ssl=1)

Karakteristik Docker Compose :

- define sebuah environment aplikasi pada `Dockerfile`
- define multi conainer app dalam file `docker-compose.yaml` atau `docker-compose.yml`
- hanya memmbutuh satu command untuk mendeploy semua aplikasi, yaitu `docker-compose up`
- menghandle dependecies dari container

Contoh `docker-compose.yml` :
```
version: "3.9"  # optional since v1.27.0
services:
  web:
    build: .
    ports:
      - "8000:5000" # port kiri bisa diakses oleh kita sdgkn kanan dikases oeh container
    volumes:
      - .:/code
      - logvolume01:/var/log
    depends_on:
      - redis
  redis:
    image: redis
volumes:
  logvolume01: {}
```



## **Volume**

Volume adalah sebuah direktori tempat menyipan data sementara dan dapat diakses ke container dalam sebuah pod. Volume disimpan di bagian sistem file host yang dikelola oleh Docker (`/var/lib/docker/volumes/` di Linux). Proses non-Docker tidak boleh memodifikasi bagian sistem file ini. Volume adalah cara terbaik untuk menyimpan data di Docker.


## **Network**

Salah satu alasan Docker Container dan service **powerful** adalah network yang bisa membuat container saling berkomunikasi bahkan dengan non-Docker workloads. Docker Container bahkan tidak perlu mengetahui apakah container di deploy dalam dalam Docker atau tidak.

## **Continuous Integration/Continuous Delivery (CI/CD)**

Continuous integration (CI) adalah pengintegrasian kode ke dalam repositori kode kemudian menjalankan pengujian secara otomatis, cepat, dan sering. CI ini dapat dilakukan dengan menggunakan perintah  `commit`.

Sementara Continuous Delivery atau Continuous Deployment (CD) adalah praktik yang dilakukan setelah proses CI selesai dan seluruh kode berhasil terintegrasi, sehingga aplikasi bisa dibangun lalu dirilis secara otomatis.

Istilah CI/CD ini biasa dipakai untuk menjalankan tes.


**Keuntungan CI/CD:**

- Mendapat feedback lebih cepat

  Dalam penggunaan CI/CD pipeline ini kode akan diuji coba secara bersamaan agar proses pengembangan perangkat lunak dapat berjalan dengan seimbang. Uji coba dilakukan dengan CI tool.

- Dapat mendeteksi bug lebih cepat

   CI/CD bekerja dengan melakukan pengujian secara otomatis, sehingga jika ada bug yang muncul pada aplikasi yang dikembangkan maka akan langsung terdeteksi oleh CI tool.

- Dapat mempercepat proses rilis

   Kode-kode terus digabungkan dan diterapkan ke dalam produk, sehingga aplikasi selalu dalam kondisi siap untuk dirilis kapan pun.

**Secara umum, CI/CD lifecycle adalah sebagai berikut :**

![](https://www.flagship.io/wp-content/uploads/ci-cd-diagram.png)

**Tools untuk CI/CD**

- Jenkins

  ![](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/Jenkins_logo.svg/1200px-Jenkins_logo.svg.png)

- Azure DevOps

  ![](https://miro.medium.com/max/800/0*cl5McDYMbuXw-fdS.png)


- GitLab CI/CD

  ![](https://camo.githubusercontent.com/07a52f72326754e7fe182adce6b572f1fc2bea794916e8c75e249d9b171fd4f6/68747470733a2f2f61626f75742e6769746c61622e636f6d2f696d616765732f63692f6769746c61622d63692d63642d6c6f676f5f32782e706e67)

- Github Action

  ![](https://avatars.githubusercontent.com/u/44036562?s=280&v=4)



## **Github action**

GitHub Actions adalah platform continuous integration dan continuous delivery (CI/CD) yang memungkinkan untuk mengotomasi alur build, test, dan deployment pipeline. Kita dapat membuat workflow yang build dan test setiap pull request ke repositori, atau menerapkan merged pull requests ke production. Untuk lebih lengkapnya dapat dilihat [di sini](https://docs.github.com/en/actions).

### **Component Github Action**

![](https://docs.github.com/assets/cb-25535/images/help/images/overview-actions-simple.png)

1. **Worklflow**

      Workflow adalah proses otomatis yang dapat dikonfigurasi yang akan menjalankan satu atau lebih pekerjaan. Workflow ditentukan oleh file YAML yang diperiksa di repositori dan akan berjalan saat dipicu oleh event di repositori, atau bisa dipicu secara manual, atau pada jadwal yang ditentukan.

2. **Event**

    Event adalah aktivitas tertentu dalam repositori yang memicu workflow berjalan. Misalnya, aktivitas dapat berasal dari GitHub saat seseorang membuat pull request, opens an issue, atau push sebuah commit ke repositori.

3. **Jobs**

    Job adalah serangkaian langkah dalam workflow yang dijalankan pada runner yang sama. Setiap langkah adalah shell script yang akan dieksekusi, atau action yang akan dijalankan.

4. **Actions**

    Action adalah aplikasi kustom untuk platform GitHub Actions yang melakukan tugas yang kompleks namun sering diulang. Gunakan tindakan untuk membantu mengurangi jumlah kode berulang yang ditulis pada file workflow. Sebuah action dapat berupa pull git repository dari GitHub, set up toolchain untuk build environment, atau set up authentication ke cloud provider.

5. **Runners**

    Runner adalah server yang menjalankan workflow saat ditrigger. Setiap runner dapat menjalankan satu job pada satu waktu. GitHub menyediakan Ubuntu Linux, Microsoft Windows, dan runner macOS untuk menjalankan workflow; setiap alur kerja dijalankan dalam virtual machine baru yang disediakan. Atau bahasa mudahnya, runner adalah sebuah aplikasi Github yang menjalankan Github logic nya.


## **Observability**

Observability adalah suatu kegiatan untuk mengukur metrics, log, dan traces.

**Kelebihan observability:**
- application performance monitoring (APM)
- DevSecOps dan site reliability engineer (SRE)
- monitoring infrastucture, cloud, dan kubernetes
- end user experience
- business analytics

**Bagaimana cara membuat sebuah sistem observable?**
- membuat logs
- membuat metric
- membuat distributed tracing
- membuat user experience

**Tantangan dalam Observability :**
- data silos
- volume, velocity, varety, complexity
- manual instrumentation and configuration
- lack of pre-production
- wasting time troubleshooting

**Tools observability**

1. [Dynatrace](https://www.dynatrace.com/)

    ![](https://www.centraldatatech.com/uploads/product/logo/Dynatrace_website.png)

2. [Lightrun](https://lightrun.com/)

    ![](https://www.stackvidhya.com/wp-content/uploads/2021/10/Lightrun-logo.svg)

3. [Lumigo](https://lumigo.io/)

    ![](https://lumigo.awsworkshop.io/images/lumigo.png)

4. [Lightstep](https://lightstep.com/)

    ![](https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_538963bea2f1be902f81195b0b990487/lightstep.png)

5. [Anodot](https://www.anodot.com/)

    ![](https://blog.anodot.com/hs-fs/hubfs/full_logo@2x-1.png?width=3780&name=full_logo@2x-1.png)

6. [Datadog](https://www.datadoghq.com/)

    ![](https://upload.wikimedia.org/wikipedia/en/thumb/7/7e/Datadog_logo.svg/1200px-Datadog_logo.svg.png)

7. [Grafana](https://grafana.com/)

    ![](https://upload.wikimedia.org/wikipedia/en/thumb/a/a1/Grafana_logo.svg/1200px-Grafana_logo.svg.png)

8. [Honeycomb](https://www.honeycomb.io/)

    ![](https://mma.prnewswire.com/media/634345/Honeycomb_Logo.jpg?p=facebook)

9. [New Relic](https://newrelic.com/)

    ![](https://play-lh.googleusercontent.com/yQj5aOQgDxKi1TuitOfVm4jyMDDbQdN6mLg_b2zZqscWBKxCvMk7atGiZBuQgKprZZY)
