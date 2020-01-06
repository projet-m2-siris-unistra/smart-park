package fr.smartpark.navigator.data

import androidx.room.Database
import androidx.room.RoomDatabase
import fr.smartpark.navigator.data.models.Tenant
import fr.smartpark.navigator.data.models.Zone

@Database(entities = [Tenant::class, Zone::class], version = 3, exportSchema = true)
abstract class AppDatabase : RoomDatabase() {
    abstract fun zoneDao(): ZoneDao
    abstract fun tenantDao(): TenantDao
}