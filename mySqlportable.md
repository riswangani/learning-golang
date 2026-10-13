# Panduan Setup MySQL Portable di Windows

Dokumentasi ini dibuat untuk kebutuhan development di laptop/komputer kantor yang membatasi hak instalasi aplikasi (UAC / Administrator blocked), sehingga MySQL dapat dijalankan secara **100% portable** di user-space.

---

## ⚡ Cara Cepat: 1x Copy-Paste Script (PowerShell)

Jika kamu sudah mendownload file **MySQL ZIP Archive** (`mysql-*-winx64.zip`) dan menyimpannya di folder `Downloads`:

Cukup buka **PowerShell** (biasa, tanpa Run as Administrator), lalu copy-paste seluruh kode di bawah ini dan tekan **Enter**:

```powershell
# ========================================================
# AUTO SETUP MYSQL PORTABLE
# ========================================================

# 1. Tentukan Path
$targetDir = "C:\Personal\mysql"
$downloads = "$env:USERPROFILE\Downloads"
$zipFile = Get-ChildItem -Path $downloads -Filter "*mysql*winx64*.zip" | Select-Object -First 1

if (-not $zipFile) {
    Write-Host "File zip MySQL tidak ditemukan di Downloads! Pastikan sudah download file .zip dari web resmi MySQL." -ForegroundColor Red
    return
}

Write-Host "Mengekstrak $($zipFile.Name) ke $targetDir..." -ForegroundColor Cyan
New-Item -ItemType Directory -Path "C:\Personal" -Force | Out-Null
tar -xf $zipFile.FullName -C "C:\Personal"

# Ganti nama folder hasil ekstrak menjadi 'mysql'
$extractedFolder = Get-ChildItem -Path "C:\Personal" -Filter "*mysql*winx64*" -Directory | Select-Object -First 1
if ($extractedFolder.FullName -ne $targetDir) {
    Rename-Item -Path $extractedFolder.FullName -NewName "mysql" -Force
}

# 2. Buat file konfigurasi my.ini
Write-Host "Membuat my.ini..." -ForegroundColor Cyan
@'
[mysqld]
basedir="C:/Personal/mysql"
datadir="C:/Personal/mysql/data"
port=3306
bind-address=127.0.0.1
max_connections=50

[client]
port=3306
'@ | Set-Content -Path "$targetDir\my.ini" -Encoding ASCII

# 3. Inisialisasi Database (Root tanpa password)
Write-Host "Menginisialisasi database pertama kali (tunggu beberapa detik)..." -ForegroundColor Cyan
& "$targetDir\bin\mysqld.exe" --defaults-file="$targetDir\my.ini" --initialize-insecure --console

# 4. Buat Script Helper (.bat)
Write-Host "Membuat file helper batch..." -ForegroundColor Cyan
@'
@echo off
echo Menjalankan MySQL Server Portable...
"%~dp0bin\mysqld.exe" --defaults-file="%~dp0my.ini" --console
'@ | Set-Content -Path "$targetDir\start-mysql.bat" -Encoding ASCII

@'
@echo off
echo Menghentikan MySQL Server...
"%~dp0bin\mysqladmin.exe" -u root shutdown
echo MySQL berhasil dimatikan.
pause
'@ | Set-Content -Path "$targetDir\stop-mysql.bat" -Encoding ASCII

@'
@echo off
"%~dp0bin\mysql.exe" -u root
'@ | Set-Content -Path "$targetDir\mysql-cli.bat" -Encoding ASCII

# 5. Daftarkan ke User PATH
Write-Host "Mendaftarkan ke User PATH..." -ForegroundColor Cyan
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$targetDir\bin*") {
    [Environment]::SetEnvironmentVariable("Path", "$userPath;$targetDir\bin", "User")
}

# 6. Buat Shortcut di Desktop
Write-Host "Membuat Shortcut di Desktop..." -ForegroundColor Cyan
$WshShell = New-Object -ComObject WScript.Shell
$desktop = [Environment]::GetFolderPath("Desktop")

$scStart = $WshShell.CreateShortcut("$desktop\Start MySQL.lnk")
$scStart.TargetPath = "$targetDir\start-mysql.bat"
$scStart.WorkingDirectory = $targetDir
$scStart.Save()

$scStop = $WshShell.CreateShortcut("$desktop\Stop MySQL.lnk")
$scStop.TargetPath = "$targetDir\stop-mysql.bat"
$scStop.WorkingDirectory = $targetDir
$scStop.Save()

Write-Host "=== SELESAI! MySQL Portable siap digunakan ===" -ForegroundColor Green
Write-Host "Shortcut 'Start MySQL' dan 'Stop MySQL' sudah ada di Desktop." -ForegroundColor Green
```

---

## 🛠️ Langkah Manual (Step-by-Step)

Jika ingin melakukan setup secara manual tahap demi tahap:

### 1. Download MySQL ZIP Archive

- Download dari: [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)
- **Jangan unduh MSI Installer**, pilih yang berformat **Windows (x86, 64-bit), ZIP Archive**.

### 2. Ekstrak ke Direktori Pilihan

Ekstrak isi zip ke folder personal, contoh: `C:\Personal\mysql`.

### 3. Buat File Konfigurasi `my.ini`

Buat file `C:\Personal\mysql\my.ini`:

```ini
[mysqld]
basedir="C:/Personal/mysql"
datadir="C:/Personal/mysql/data"
port=3306
bind-address=127.0.0.1
max_connections=50

[client]
port=3306
```

### 4. Inisialisasi Database

Jalankan perintah ini di PowerShell:

```powershell
& "C:\Personal\mysql\bin\mysqld.exe" --defaults-file="C:\Personal\mysql\my.ini" --initialize-insecure --console
```

> Parameter `--initialize-insecure` membuat akun `root` tanpa password sehingga siap pakai.

### 5. Tambahkan ke User PATH

```powershell
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
[Environment]::SetEnvironmentVariable("Path", "$userPath;C:\Personal\mysql\bin", "User")
```

---

## 🚀 Penggunaan Sehari-hari

### 1. Menjalankan Server MySQL

- **Via Shortcut Desktop:** Double-click **`Start MySQL`**.
- **Via Terminal (Foreground):**
  ```powershell
  mysqld --defaults-file="C:\Personal\mysql\my.ini" --console
  ```
- **Via Terminal (Background / Tanpa Jendela):**
  ```powershell
  Start-Process "mysqld" -ArgumentList '--defaults-file="C:\Personal\mysql\my.ini"' -WindowStyle Hidden
  ```

### 2. Konek ke MySQL CLI

```powershell
mysql -u root
```

_(Langsung tekan Enter, tanpa password)_.

### 3. Mematikan Server MySQL

- **Via Shortcut Desktop:** Double-click **`Stop MySQL`**.
- **Via Terminal:**
  ```powershell
  mysqladmin -u root shutdown
  ```

### 4. Cek Status MySQL

```powershell
mysqladmin -u root ping
# Output: mysqld is alive
```

---

## 🔌 Konfigurasi Koneksi (DBeaver / DataGrip / Backend)

- **Host:** `localhost` atau `127.0.0.1`
- **Port:** `3306`
- **Username:** `root`
- **Password:** _(Kosong / Blank)_

### Contoh Koneksi di Golang (`database/sql` & `go-sql-driver/mysql`):

```go
db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nama_database?parseTime=true")
```
