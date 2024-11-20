create table m_poli(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
poli varchar(255),
maksimal_pasien int,
created_at date,
updated_at date
)

create table m_seqno_antrian_poli(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
tanggal date,
kode_poli varchar(30),
kode_dokter varchar(30),
seqno int,
status varchar(30), --dilayani, tidak ada, pending
created_at date,
updated_at date
)

create table m_kategori_petugas(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
deskripsi varchar(50),
created_at date,
updated_at date
) 

create table m_petugas(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
nama varchar(255),
gelar varchar(50),
kode_kategori_petugas varchar(30),
password varchar(255),
created_at date,
updated_at date
)

create table m_pasien(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
kode varchar(30),
nama varchar(255),
jenis_pasien varchar(30), --bpjs, umum, asuransi
created_at date,
updated_at date
)

create table pendaftaran(
id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
nomor_pendaftaran varchar(30),
nomor_antrian varchar(30),
nomor_antrian_poli int,
kode_pasien varchar(30),
nama_pasien varchar(255),
kode_dokter varchar(30),
nama_dokter varchar(255),
kode_poli varchar(30),
nama_poli varchar(50),
kode_jenis_pasien varchar(30), --bpjs, umum, asuransi
created_at date,
updated_at date
)

insert into m_poli(kode,poli,maksimal_pasien)
select 'P01','Poli Umum',30
union all
select 'P02','Poli Bpjs',30
union all
select 'P03','Poli Penyakit dalam',30

insert into m_kategori_petugas(kode, deskripsi)
select 'DO','Dokter'
union all
select 'LO','Loket'
union all
select 'PR','Perawat'
union all
select 'PL','Petugas Poli'

insert into m_petugas(kode, nama,gelar,kode_kategori_petugas,password)
select 'D0001','Irwansyah','Spog','DO','password'
union all
select 'D0002','Yusuf','Spd','DO','password'
union all
select 'L0001','Anhari',null,'LO','password'
union all
select 'L0002','Nismawati',null,'LO','password'