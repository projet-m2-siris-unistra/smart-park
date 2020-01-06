package fr.smartpark.navigator.inject

import android.content.Context
import androidx.room.Room
import dagger.Module
import dagger.Provides
import fr.smartpark.navigator.data.AppDatabase
import fr.smartpark.navigator.data.ParkingZoneRepository
import fr.smartpark.navigator.utilities.DATABASE_NAME
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
            .build()

    @JvmStatic
    @Singleton
    @Provides
    fun provideIoDispatcher() = Dispatchers.IO
}
