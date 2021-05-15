#!/bin/sh

installReqPackages(){
  apt-get -y install "$1"
  # shellcheck disable=SC2181
  if [ "$?" -ne 0 ];then
    #res="Failed"
    echo "Failed to install $1 package.Hence, exiting installation."
    echo "Please try to install manually/delete and try again."
    exit 1
  fi
}
getBinaryPath(){
  binaryPath=$(command -v "$1")
}
getPythonPkgMgr(){
  getBinaryPath 'python'
  if [ "$binaryPath" = '' ];then
    getBinaryPath "python2"
    pythonPkgMgr='pip2'
    if [ "$binaryPath" = '' ];then
      getBinaryPath "python3"
      pythonPkgMgr='pip3'
    fi
  else
    pythonPkgMgr='pip'
  fi
}
installPythonModules(){
  pip3 install "$1"
}

createKeyboardShortCut(){
  echo "pending"
}
createBootupfile(){
  cp /opt/linux-personal-assistance/command-handlers/greet_user /etc/init.d/
  chmod 766 /etc/init.d/greet_user
  ln -s /etc/init.d/greet_user  /etc/rc5.d/S01linux-personal-assistance
  chmod 777 /etc/rc5.d/S01linux-personal-assistance
}
createFolderStructure(){
  mkdir -p /opt/linux-personal-assistance/linux-assistance
}


#getPythonPkgMgr
echo "================================="
echo "Installing supporting packages"
installReqPackages "python3-pip"
#getBinaryPath "$pythonPkgMgr"
pythonPkgMgr='pip3'
installPythonModules 'SpeechRecognition'
if [ "$pythonPkgMgr" = 'pip3' ];then
  installReqPackages 'python3-pyaudio'
else
  installReqPackages 'python-pyaudio'
fi
installReqPackages 'mplayer'
