package fr.smartpark.navigator.inject

import android.content.Context
import androidx.room.Room
import dagger.Module
import dagger.Provides
import fr.smartpark.navigator.api.ApiService
import fr.smartpark.navigator.api.ZoneRemoteDataSource
import fr.smartpark.navigator.data.AppDatabase
import fr.smartpark.navigator.data.ZoneDao
import fr.smartpark.navigator.data.ZoneRepository
import fr.smartpark.navigator.utilities.DATABASE_NAME
import kotlinx.coroutines.Dispatchers
import javax.inject.Singleton

@Module
object AppModule {
    @JvmStatic
    @Singleton
    @Provides
    fun provideZoneRemoteDataSource(service: ApiService) =
        ZoneRemoteDataSource(service)

    @JvmStatic
    @Singleton
    @Provides
    fun provideZoneDao(database: AppDatabase) =
        database.zonesDao()

    @JvmStatic
    @Singleton
    @Provides
    fun provideZoneRepository(dao: ZoneDao, remoteDataSource: ZoneRemoteDataSource) =
        ZoneRepository(dao, remoteDataSource)

    @JvmStatic
    @Singleton
    @Provides
    fun provideDatabase(context: Context) =
        Room.databaseBuilder(context.applicationContext, AppDatabase::class.java, DATABASE_NAME)
            .fallbackToDestructiveMigration()
            .build()

    @JvmStatic
    @Singleton
    @Provides
    fun provideIoDispatcher() = Dispatchers.IO
}
