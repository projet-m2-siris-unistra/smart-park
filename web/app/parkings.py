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



# Instance of a tenant
class TenantManagement:

    def __init__(self, tenant_id):
        self.id = tenant_id
        self.zones = []
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
                item['name'],
                item['type'],
                '#' + item['color'],
                item['geo']
            )
            self.zones.append(obj)


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


    def toJson(self):
        return {
            "id" : self.id,
            "name" : self.name,
            "nb_total_spots" : self.getNbTotalSpots(),
            "nb_taken_spots" : self.getNbTakenSpots(),
            "desc" : self.desc,
            "type" : self.type,
            "color" : self.color,
            "coordinates" : self.polygon
        }

    # Getter / Setter #

    def getNbTotalSpots(self):
        # calculation from DB
        return 321


    def getNbTakenSpots(self):
        # calculation from DB
        return 123


    async def getSpotList(self):
        # requesting all spots belonging to this zone
        # loop for parsing all spots
        response = await Request.getSpots(self.id)
        data = js.loads(response)

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

    # Data requests # 
    
    # When displaying marker, we customi
    #def coordinatesGeoJson(self):
    #    return {
    #        'type': 'Feature',
    #        'geometry': {
    #            'type': 'Point',
    #            'coordinates': self.coordinates
    #        }
    #    }


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

        self.battery = data['battery']
        self.state = data['state']


    # returns objects attributs wrapped into Json
    def toJson(self):
        return {
            "id" : self.id,
            "state" : self.state,
            "battery" : self.battery
        }