import json as js
import datetime

from app.bus import Request

class Tooling:

    # convert elements of list into Json
    @staticmethod
    def jsonList(arg):
        liste = []
        for item in arg:
            liste.append(js.dumps(item.toJson()))
        return liste


    # format #ffffff color into right format FFFFFF
    @staticmethod
    def formatColor(arg):
        return arg[1:].upper()



# Instance of a tenant
class TenantManagement:

    def __init__(self, tenant_id):
        self.id = tenant_id
        self.zones = []
        self.devices = []
        # default data to see if DB requests are doing well
        self.name = "NOT UPDATED"
        self.coordinates = [7.9726, 49.0310] # Altenstadt (FR,67)


    async def init(self, tenant_id):
        response = await Request.getTenant(tenant_id)
        data = js.loads(response)
        self.name = data['name']
        self.coordinates = data['geo']


    # Get the list of all the zones from this tenant
    async def setZones(self):
        reponse = await Request.getZones(self.id)
        data = js.loads(reponse)

        for item in data:
            obj = ZoneManagement(item['zone_id'])
            obj.staticInit(
                name=item['name'],
                type=item['type'],
                color='#' + item['color'],
                polygon=item['geo']
            )
            self.zones.append(obj)


    # Get the list of all the NOT ASSIGNED devices of this tenant
    async def setDevices(self):
        reponse = await Request.getDevices(self.id)
        data = js.loads(reponse)

        for item in data:
            obj = DeviceManagement(item['device_id'])
            obj.staticInit(
                eui=item['device_eui'],
                battery=item['battery'],
                state=item['state']
            )
            self.devices.append(obj)


    # This function returns the list of all not assigned devices
    # return a list of form: ['device_id1':'eui1', 'device_id2':'eui2']
    async def getNotAssignedDevices(self):
        response = await Request.getNotAssignedDevices(self.id)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR
            
        data = js.loads(response)
        print("data=", data)

        dict = []
        for item in data:
            dict[item['device_id']] = item['eui']
        
        print("dict=", dict)
        return dict
        

    def getTotalSpots(self):
        count = 0
        if self.zones == []:
            print("WARNING: tenant.zones empty")
            return -1
        for zone in self.zones:
            count += zone.getNbTotalSpots()
        return count


    def getTakenSpots(self):
        count = 0
        if self.zones == []:
            print("WARNING: tenant.zones empty")
            return -1
        for zone in self.zones:
            count += zone.getNbTakenSpots()
        return count



# Instance of a zone
class ZoneManagement:

    # ID for non persistent objects (I.e. not in DB)
    notAssigned = -1

    def __init__(self, zone_id):
        self.id = zone_id
        # some default data to reveal further failed init
        self.desc = "Parking description"
        self.spots = []


    def staticInit(self, name, type, color, polygon):
        self.name = name
        self.type = type
        self.color = color
        self.polygon = polygon


    async def init(self, zone_id):
        response = await Request.getZone(zone_id)
        data = js.loads(response)
        self.name = data['name']
        self.type = data['type']
        self.color = '#' + data['color']
        self.polygon = data['geo']


    async def create(self, tenant_id):
        print("ZoneInstance.create called")
        response = await Request.createZone(
            tenant_id,
            self.name,
            self.type,
            self.color,
            self.polygon
        )
        # Checking response ?


    def toJson(self):
        if not (self.spots is None):
            spotJson = self.spotsToJson()
        return {
            "id" : self.id,
            "name" : self.name,
            "nb_total_spots" : self.getNbTotalSpots(),
            "nb_taken_spots" : self.getNbTakenSpots(),
            "desc" : self.desc,
            "type" : self.type,
            "color" : self.color,
            "coordinates" : self.polygon,
            "spots" : spotJson
        }
    
    def spotsToJson(self):
        listJson = []
        for spot in self.spots:
            listJson.append(spot.toJson())
        return listJson

    # Getter / Setter #

    def getNbTotalSpots(self):
        # calculation from DB
        return 321


    def getNbTakenSpots(self):
        # calculation from DB
        return 123


    async def setSpots(self):
        # requesting all spots belonging to this zone
        # loop for parsing all spots
        response = await Request.getSpots(self.id)
        data = js.loads(response)

        if data is not None:
            for item in data:
                obj = SpotManagement(item['place_id'])
                obj.staticInit(
                    item['place_id'],
                    item['geo'],
                    item['type'],
                    item['device_id']
                )
                await obj.setDevice(item['device_id'])
                self.spots.append(obj)

     # Statistics #

    def getDailyStats(self):
        stats = {
            'stats_type':'Quotidienne',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    
    
    def getWeeklyStats(self):
        stats = {
            'stats_type':'Hebdomadaire',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats


    def getMonthlyStats(self):
        stats = {
            'stats_type':'Mensuelle',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats


    def getAnnualStats(self):
        stats = {
            'stats_type':'Annuelle',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    

    def getAllStats(self):
        stats = []
        stats.append(self.getDailyStats())
        stats.append(self.getWeeklyStats())
        stats.append(self.getMonthlyStats())
        stats.append(self.getAnnualStats())
        return stats



# Instance of a parking spot 
class SpotManagement:

    def __init__(self, spot_id):
        self.id = spot_id
        self.name = "default"
        self.coordinates = [7.9726, 49.0310] # Altenstadt (FR,67)


    async def init(self, spot_id):
        response = await Request.getSpot(spot_id)
        data = js.loads(response)
        self.coordinates = data['geo']
        self.type = data['type']
        self.name = "PLACE-" + str(spot_id)
        device_id = data['device_id']


    def staticInit(self, spot_id, coordinates, type, device_id):
        self.coordinates = coordinates
        self.type = type
        self.name = "PLACE-" + str(spot_id)


    async def setDevice(self, device_id):
        deviceInstance = DeviceManagement(device_id)
        await deviceInstance.init(device_id)
        self.device = deviceInstance


    def toJson(self):
        return {
            "id" : self.id,
            "name" : self.name,
            "type" : self.type,
            "coordinates" : self.coordinates,
            "device" : self.device.toJson()
        }

    # Statistics #

    def getDailyStats(self):
        stats = {
            'stats_type':'Daily',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    

    def getWeeklyStats(self):
        stats = {
            'stats_type':'Weekly',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats


    def getMonthlyStats(self):
        stats = {
            'stats_type':'Monthly',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats


    def getAnnualStats(self):
        stats = {
            'stats_type':'Annual',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats


    def getAllStats(self):
        stats = []
        stats.append(self.getDailyStats())
        stats.append(self.getWeeklyStats())
        stats.append(self.getMonthlyStats())
        stats.append(self.getAnnualStats())
        return stats



# instance of a device
class DeviceManagement:

    def __init__(self, device_id):
        self.id = device_id
        # default values to see if request failed
        self.battery = -1


    async def init(self, device_id):
        response = await Request.getDevice(device_id)
        data = js.loads(response)

        self.eui = data['device_eui']
        self.battery = data['battery']
        self.state = data['state']


    # Manual initialization of a device object
    def staticInit(self, eui, battery, state):
        self.eui = eui
        self.battery = battery
        self.state = state


    # returns objects attributs wrapped into Json
    def toJson(self):
        return {
            "id" : self.id,
            "eui" : self.eui,
            "state" : self.state,
            "battery" : self.battery
        }