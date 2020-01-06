package fr.smartpark.navigator


import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.viewModels
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.observe
import androidx.navigation.fragment.navArgs
import androidx.transition.TransitionInflater
import dagger.android.support.DaggerFragment
import fr.smartpark.navigator.databinding.FragmentParkingZoneDetailBinding
import fr.smartpark.navigator.viewmodels.ParkingZoneDetailViewModel
import javax.inject.Inject

class ParkingZoneDetailFragment : DaggerFragment() {
    @Inject
    lateinit var viewModelFactory: ViewModelProvider.Factory
    private val viewModel by viewModels<ParkingZoneDetailViewModel> { viewModelFactory }

    private lateinit var binding: FragmentParkingZoneDetailBinding

    private val args: ParkingZoneDetailFragmentArgs by navArgs()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        sharedElementEnterTransition =
            TransitionInflater.from(context).inflateTransition(R.transition.parking_detail)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentParkingZoneDetailBinding.inflate(inflater, container, false)
        viewModel.start(args.parkingZone.zoneId)

        viewModel.zone.observe(viewLifecycleOwner) { zone ->
            binding.zone = zone
        }

        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding.apply {
            // Fallback to zone passed in args
            zone = viewModel.zone.value ?: args.parkingZone
            executePendingBindings()
        }
    }
}
