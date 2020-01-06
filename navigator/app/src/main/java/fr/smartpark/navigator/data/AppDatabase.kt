package fr.smartpark.navigator.data

import androidx.room.Database
import androidx.room.RoomDatabase
import fr.smartpark.navigator.data.models.Zone

@Database(entities = [Zone::class], version = 2, exportSchema = true)
abstract class AppDatabase : RoomDatabase() {
    abstract fun zonesDao(): ZoneDao
}