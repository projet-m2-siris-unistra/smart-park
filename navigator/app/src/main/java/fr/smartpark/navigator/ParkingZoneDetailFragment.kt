package fr.smartpark.navigator


import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.navArgs
import fr.smartpark.navigator.databinding.FragmentParkingZoneDetailBinding

class ParkingZoneDetailFragment : Fragment() {
    private lateinit var binding: FragmentParkingZoneDetailBinding

    private val args: ParkingZoneDetailFragmentArgs by navArgs()

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // TODO: rebind to a LiveData<ParkingZone>
        binding = FragmentParkingZoneDetailBinding.inflate(inflater, container, false)
        binding.zone = args.parkingZone
        return binding.root
    }

}
