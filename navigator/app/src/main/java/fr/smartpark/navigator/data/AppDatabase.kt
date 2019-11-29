package fr.smartpark.navigator.data

import android.content.Context
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase
import androidx.sqlite.db.SupportSQLiteDatabase
import androidx.work.OneTimeWorkRequestBuilder
import androidx.work.WorkManager
import fr.smartpark.navigator.utilities.DATABASE_NAME
import fr.smartpark.navigator.workers.SeedDatabaseWorker

@Database(entities = [ParkingZone::class], version = 1, exportSchema = true)
abstract class AppDatabase : RoomDatabase() {
    abstract fun zonesDao(): ParkingZoneDao

    companion object {
        @Volatile private var instance: AppDatabase? = null

        operator fun invoke(context: Context) =
            instance ?: synchronized(this) {
                instance ?: buildDatabase(context).also { instance = it }
            }

        private fun buildDatabase(context: Context) =
            Room.databaseBuilder(context, AppDatabase::class.java, DATABASE_NAME)
                .addCallback(object : RoomDatabase.Callback() {
                    override fun onCreate(db: SupportSQLiteDatabase) {
                        super.onCreate(db)
                        val request = OneTimeWorkRequestBuilder<SeedDatabaseWorker>().build()
                        WorkManager.getInstance(context).enqueue(request)
                    }
                })
                .build()
    }
}