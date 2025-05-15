@echo off
echo Running Go Web App...

REM Copy required directories to bin if they don't exist
if not exist "bin\templates" (
    echo Copying templates directory...
    xcopy /E /I /Y templates bin\templates
)

if not exist "bin\static" (
    echo Copying static directory...
    xcopy /E /I /Y static bin\static
)

REM Run the application
cd bin
server.exe
pause
