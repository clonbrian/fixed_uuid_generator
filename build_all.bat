@echo off
setlocal

echo ========================================
echo   UUID Generator - Build All
echo ========================================
echo.

REM Remove old outputs
if exist dist\uuid_generator.exe del dist\uuid_generator.exe
if exist dist\uuid_generator_mac del dist\uuid_generator_mac
if exist dist\uuid_generator_mac_intel del dist\uuid_generator_mac_intel

REM Create dist folder
if not exist dist mkdir dist

REM Build Windows
echo [1/3] Building Windows: dist\uuid_generator.exe ...
go build -o dist\uuid_generator.exe .
if errorlevel 1 goto build_failed
echo       OK.
echo.

REM Build Mac Apple Silicon
echo [2/3] Building Mac Apple Silicon: dist\uuid_generator_mac ...
set GOOS=darwin
set GOARCH=arm64
go build -o dist\uuid_generator_mac .
if errorlevel 1 goto build_failed
echo       OK.
echo.

REM Build Mac Intel
echo [3/3] Building Mac Intel: dist\uuid_generator_mac_intel ...
set GOARCH=amd64
go build -o dist\uuid_generator_mac_intel .
if errorlevel 1 goto build_failed
echo       OK.
echo.

REM Done
echo ========================================
echo   Build OK.
echo ========================================
echo.
echo Output in dist:
echo   uuid_generator.exe       - Windows
echo   uuid_generator_mac       - Mac Apple Silicon
echo   uuid_generator_mac_intel - Mac Intel
echo.
goto end

:build_failed
echo.
echo ========================================
echo   Build FAILED. Check error above.
echo ========================================
exit /b 1

:end
