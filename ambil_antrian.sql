CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table m_jenis_pasien(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
deskripsi varchar(255),
created_at date,
updated_at date
)

create table m_antrian_kategori(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
deskripsi varchar(255),
kode_jenis_pasien varchar(30),
created_at date,
updated_at date
)

create table m_seqno_antrian(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
tanggal date,
kode_antrian_kategori varchar(30),
seqno int,
status varchar(30), --dilayani, tidak ada, pending
panggil int,
aktif int,
kode_loket varchar(10),
created_at TIMESTAMP,
updated_at TIMESTAMP
)

select to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS') ,* from m_seqno_antrian
ALTER TABLE m_seqno_antrian
ALTER COLUMN updated_at TYPE TIMESTAMP
USING updated_at::TIMESTAMP;

SET TIMEZONE = 'Asia/Jakarta'; -- atau zona waktu lain yang relevan
UPDATE m_seqno_antrian SET updated_at = NOW();



create table m_status(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(10),
deskripsi varchar(30),
created_at date,
updated_at date
)


create table m_perusahaan(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
nama varchar(255),
image varchar(255),
created_at date,
updated_at date
)

create table m_loket(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
deskripsi varchar(255),
aktif int,
created_at date,
updated_at date
)

create table m_loket_log(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode_loket varchar(10),
kode_petugas varchar(20),
created_at date,
updated_at date
)

insert into m_jenis_pasien(kode, deskripsi)
select 'bpjs','BPJS'
union all
select 'umum','Umum'
union all
select 'asuransi','Asuransi'

insert into m_antrian_kategori(kode, deskripsi, kode_jenis_pasien)
select 'A','Pasien Umum','umum'
union all
select 'B','Pasien BPJS','bpjs'
union all
select 'C','Pasien Asuransi','asuransi'
union all
select 'JM','Pasien BPJS','bpjs'

insert into m_status(kode,deskripsi)
select '1','Ambil Antrian'
union all
select '2','Pendaftaran'
union all
select '3','TTV'
union all
select '4','Dilayani'
union all
select '0','Tidak Jadi'


insert into m_perusahaan(kode,nama)
select 'A01','RSUD. Cileungsi'


insert into m_loket(kode, deskripsi,aktif)
select '1','Loket 1',1
union all
select '2','Loket 2',1
union all
select '3','Loket 3',1
union all 
select '4','Loket 4',1
union all
select '5','Loket 5',1
union all 
select '6','Loket 6',1

insert into m_seqno_antrian (tanggal, kode_antrian_kategori,seqno,status)
select now(),'A',1,1
union all
select now(),'A',2,1
union all
select now(),'A',3,1
union all
select now(),'A',4,1
union all
select now(),'B',1,1
union all
select now(),'B',2,1
union all
select now(),'B',3,1
union all
select now(),'C',1,1
union all
select now(),'C',2,1
union all
select now(),'JM',1,1

//nomor panggil

select * from m_seqno_antrian msa 
update m_seqno_antrian  set 
panggil = 1,
kode_loket = 2,
seqno = 32
where id = '12ce5899-5b12-4d27-bca1-db7103eed185'
select * from m_loket

--1 . Panggilan
      --kolom panggil dijadikan 0 semua lalu update jadi 1 sesuai kode_antrian+seqno
      --kolom kode loket di update sesuai loket yang manggil berdasarkan kode_antrian+seqno
      --aktif dijadikan 0 untuk berdasarakan loket yang di pilih lalu dijadikan 1 berdasarakan kode_antrian+seqno
       
select * from m_seqno_antrian msa 
update m_seqno_antrian set 
panggil = 1
where id  IN('12ce5899-5b12-4d27-bca1-db7103eed185')


update m_seqno_antrian set 
tanggal = now()

 UPDATE "m_seqno_antrian" SET "tanggal"='2024-11-20 00:00:00',"kode_antrian_kategori"='C',"seqno"=2,"status"=1,"panggil"=1,"aktif"=0,"kode_loket"='3' WHERE "id" = '7e920e99-6e9d-4d80-8f92-d819af2022b8'
