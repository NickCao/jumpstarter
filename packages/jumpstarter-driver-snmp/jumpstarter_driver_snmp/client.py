from dataclasses import dataclass

from jumpstarter_driver_power.client import PowerClient


@dataclass(kw_only=True)
class SNMPServerClient(PowerClient):
    """Client interface for SNMP Power Control"""

    pass
