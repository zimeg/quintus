{
  description = "the teller of time";
  inputs = {
    cicero = ../cicero;
  };
  outputs = { self, ... }: {
    networking = {
      hostName = "tullius";
      networkmanager = {
        enable = true;
      };
      firewall = {
        enable = true;
        allowedUDPPorts = [ 123 ];
      };
    };
    systemd.services = {
      cicero = {
        enable = true;
        after = [ "network.target" ];
        wantedBy = [ "multi-user.target" ];
        serviceConfig = {
          ExecStart = "${self.inputs.cicero.packages.default}/bin/cicero";
          Restart = "always";
          RestartSec = 2;
        };
      };
    };
    time = {
      timeZone = "Etc/UTC";
    };
    users.users.default = {
      isNormalUser = true;
      name = "qts";
      password = "placeholder";
      extraGroups = [ "networkmanager" "wheel" ];
      linger = true;
    };
  };
}
