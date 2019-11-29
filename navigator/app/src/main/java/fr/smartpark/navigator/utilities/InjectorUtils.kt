package fr.smartpark.navigator.utilities

import android.content.Context
import fr.smartpark.navigator.data.AppDatabase
import fr.smartpark.navigator.data.ParkingZoneRepository
import fr.smartpark.navigator.viewmodels.ParkingZoneListViewModelFactory

object InjectorUtils {
    private fun provideParkingZoneRepository(context: Context): ParkingZoneRepository {
        return ParkingZoneRepository(AppDatabase(context).zonesDao())
    }

    fun provideParkingZoneListViewModelFactory(context: Context): ParkingZoneListViewModelFactory {
        return ParkingZoneListViewModelFactory(provideParkingZoneRepository(context))
    }
}