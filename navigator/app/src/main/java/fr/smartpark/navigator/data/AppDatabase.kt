package fr.smartpark.navigator.data

import androidx.room.Database
import androidx.room.RoomDatabase

@Database(entities = [ParkingZone::class], version = 1, exportSchema = true)
abstract class AppDatabase : RoomDatabase() {
    abstract fun zonesDao(): ParkingZoneDao
}