package fr.smartpark.navigator.data

import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.TypeConverters
import fr.smartpark.navigator.data.models.Tenant
import fr.smartpark.navigator.data.models.Zone

@Database(entities = [Tenant::class, Zone::class], version = 4, exportSchema = true)
@TypeConverters(Converters::class)
abstract class AppDatabase : RoomDatabase() {
    abstract fun zoneDao(): ZoneDao
    abstract fun tenantDao(): TenantDao
}