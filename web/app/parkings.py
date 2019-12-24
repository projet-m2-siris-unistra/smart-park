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



"""
NOTE
Data from the backend is formated in JSON. When a list is returned:
{
    'count' : 10,       // the total number of existing entries
    'data' : {...}      // the actual data
}
"""



# Instance of a tenant
class TenantManagement:

    def __init__(self, tenant_id):
        self.id = tenant_id
        self.zones = []
        self.devices = []
        self.notAssignedDevices = []
        # default data to see if DB requests are doing well
        self.name = "NOT UPDATED"
        self.coordinates = [7.9726, 49.0310] # Altenstadt (FR,67)


    async def init(self, tenant_id):
        response = await Request.getTenant(tenant_id)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

        data = js.loads(response)
        self.name = data['name']
        self.coordinates = data['geo']


    # Get the list of all the zones from this tenant
    async def setZones(self, page=1, pagesize=20):
        response = await Request.getZones(self.id, page, pagesize)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

        data = js.loads(response)
        self.zonesCount = data['count']
        self.zones.clear() # In case of...
        
        if data['data'] is not None:
            for item in data['data']:
                obj = ZoneManagement(item['zone_id'])
                obj.staticInit(
                    name=item['name'],
                    type=item['type'],
                    color='#' + item['color'],
                    polygon=item['geo']
                )
                self.zones.append(obj)


    # Get the list of all devices of this tenant
    async def setDevices(self, page=1, pagesize=20):
        response = await Request.getDevices(self.id, page, pagesize)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

        data = js.loads(response)
        #print("data", data)
        self.devicesCount = data['count']
        self.devices.clear() # In cas of...

        for item in data['data']:
            obj = DeviceManagement(item['device_id'])
            obj.staticInit(
                eui=item['device_eui'],
                battery=item['battery'],
                state=item['state']
            )
            self.devices.append(obj)


    # Get the list of all the NOT ASSIGNED devices of this tenant
    async def setNotAssignedDevices(self, page=1, pagesize=20):
        response = await Request.getNotAssignedDevices(self.id, page, pagesize)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

        data = js.loads(response)
        #print("data", data)
        self.devicesNotAssignedCount = data['count']
        self.notAssignedDevices.clear() # In cas of...

        if data['data'] is not None:
            for item in data['data']:
                obj = DeviceManagement(item['device_id'])
                obj.staticInit(
                    eui=item['device_eui'],
                    battery=item['battery'],
                    state=item['state']
                )
                self.notAssignedDevices.append(obj)


    # This function returns the list of all not assigned devices
    # It should be reduced later by using the function above
    # return a list of form: ['device_id1':'eui1', 'device_id2':'eui2']
    async def getNotAssignedDevices(self):
        response = await Request.getNotAssignedDevices(self.id, page=1, pagesize=50)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR
            
        data = js.loads(response)
        self.devicesNotAssignedCount = data['count']
        print("data", data)

        devList = []
        if data['data'] is not None:
            for item in data['data']:
                devList.append((item['device_id'], item['device_eui']))
        
        return devList
        

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
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

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

        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR


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
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

        data = js.loads(response)
        self.spotsCount = data['count']

        if data['data'] is not None:
            for item in data['data']:
                obj = SpotManagement(item['place_id'])
                obj.staticInit(
                    item['place_id'],
                    item['geo'],
                    item['type'],
                    item['device_id']
                )
                await obj.setDevice()
                self.spots.append(obj)


    # Call this method to delete this zone from database
    async def delete(self):
        response = "waiting for backend..."


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
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR
        
        data = js.loads(response)
        self.coordinates = data['geo']
        self.type = data['type']
        self.name = "PLACE-" + str(spot_id)
        self.device_id = data['device_id']


    def staticInit(self, spot_id, coordinates, type, device_id):
        self.coordinates = coordinates
        self.type = type
        self.name = "PLACE-" + str(spot_id)
        self.device_id = device_id


    async def setDevice(self):
        deviceInstance = DeviceManagement(self.device_id)
        response = await deviceInstance.init(self.device_id)
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR
        
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
        if response == Request.REQ_ERROR:
            return Request.REQ_ERROR

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