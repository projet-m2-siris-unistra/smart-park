package fr.smartpark.navigator.inject

import android.content.Context
import androidx.room.Room
import androidx.room.RoomDatabase
import androidx.sqlite.db.SupportSQLiteDatabase
import androidx.work.OneTimeWorkRequestBuilder
import androidx.work.WorkManager
import dagger.Module
import dagger.Provides
import fr.smartpark.navigator.data.AppDatabase
import fr.smartpark.navigator.data.ParkingZoneRepository
import fr.smartpark.navigator.utilities.DATABASE_NAME
import fr.smartpark.navigator.workers.SeedDatabaseWorker
import kotlinx.coroutines.Dispatchers
import javax.inject.Singleton

@Module
object AppModule {
    @JvmStatic
    @Singleton
    @Provides
    fun provideParkingZoneRepository(database: AppDatabase) =
        ParkingZoneRepository(database.zonesDao())

    @JvmStatic
    @Singleton
    @Provides
    fun provideDatabase(context: Context) =
        Room.databaseBuilder(context.applicationContext, AppDatabase::class.java, DATABASE_NAME)
            .addCallback(object : RoomDatabase.Callback() {
                override fun onCreate(db: SupportSQLiteDatabase) {
                    super.onCreate(db)
                    val request = OneTimeWorkRequestBuilder<SeedDatabaseWorker>().build()
                    WorkManager.getInstance(context).enqueue(request)
                }
            })
            .build()

    @JvmStatic
    @Singleton
    @Provides
    fun provideIoDispatcher() = Dispatchers.IO
}
