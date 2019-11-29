package fr.smartpark.navigator.data

class ParkingZoneRepository private constructor(private val zoneDao: ParkingZoneDao) {
    fun getZones() = zoneDao.getZones()

    companion object {
        @Volatile
        private var instance: ParkingZoneRepository? = null

        operator fun invoke(zoneDao: ParkingZoneDao) =
            instance ?: synchronized(this) {
                instance ?: ParkingZoneRepository(zoneDao).also { instance = it }
            }
    }
}