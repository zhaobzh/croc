package crypt

import "testing"

func BenchmarkEncryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt([]byte(`

		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse laoreet justo quis augue vehicula ornare. Nullam nec risus volutpat nulla dictum consectetur. Phasellus porttitor, justo non tincidunt finibus, massa justo iaculis urna, eget mattis nulla libero vitae risus. Vestibulum vehicula nunc id dignissim rutrum. Cras varius ac nulla a imperdiet. Sed finibus, libero in tempor hendrerit, turpis erat faucibus nisl, a consectetur massa mi eget mi. Vestibulum pulvinar lorem id ipsum elementum auctor.
		
		Morbi at odio a eros eleifend faucibus. Sed at tempor urna, in interdum neque. Curabitur condimentum rhoncus orci, vel vulputate risus ultricies efficitur. Curabitur eu vehicula ligula. Curabitur suscipit ex vitae nunc faucibus vehicula. Mauris sed dictum mauris. Vivamus nec dui at urna porttitor suscipit. Integer ut eros finibus orci consectetur hendrerit id quis dolor. Proin dapibus orci quis massa viverra finibus. Sed quis ligula neque.
		
		Aenean et fringilla nulla. In et venenatis massa, vel feugiat diam. Sed ornare felis nec egestas suscipit. Sed a ultricies sapien. Aliquam ex leo, tincidunt faucibus neque non, vehicula commodo velit. Vestibulum molestie efficitur velit in vestibulum. Aliquam eget leo felis. Etiam pharetra vulputate egestas. Cras gravida nibh eu sollicitudin facilisis. Nulla facilisi. Nulla tristique arcu vitae arcu pulvinar feugiat. Suspendisse finibus a urna a cursus. Donec bibendum sodales nunc eget tincidunt. Pellentesque blandit ac nunc at consectetur.
		
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce congue nibh nisl, et porttitor mauris tempor eget. Vivamus sit amet sagittis ligula. Quisque aliquam orci non odio egestas, vel semper felis efficitur. Vivamus non leo lacus. Etiam non ligula eget erat porta laoreet id id ligula. Proin imperdiet erat id efficitur vehicula. Suspendisse potenti. Morbi ornare finibus metus, eu pretium leo tristique sit amet. Praesent placerat elit quis porttitor rhoncus. Vestibulum consectetur turpis sed lacus placerat, vel laoreet est interdum. Proin id quam ut risus tempor hendrerit nec eu lacus. Duis vel aliquam ex.
		
		Sed facilisis in ex vitae pellentesque. Donec tempor lobortis dui. Praesent sagittis, elit ac dictum hendrerit, eros risus auctor erat, in pretium nulla mi et felis. Nullam imperdiet erat id erat rutrum fermentum. Praesent ultrices, diam non efficitur gravida, sem quam fermentum est, vitae tristique libero velit eget turpis. Nunc sem risus, venenatis nec suscipit quis, accumsan eu nunc. Fusce fringilla sit amet purus vel ornare. Donec nulla dolor, gravida ac metus nec, tincidunt feugiat sem. Fusce semper varius nunc.
		
		Nulla facilisi. Quisque in euismod ex, ac sollicitudin dolor. Aliquam erat volutpat. Etiam nisl sem, posuere et pellentesque fermentum, pretium non lorem. Nunc blandit nisl at leo interdum gravida. Proin ullamcorper ultrices dictum. Pellentesque non ligula magna. Sed justo nibh, finibus vitae malesuada ultricies, molestie ac ante. Suspendisse maximus congue viverra.
		
		Donec a est eu arcu tristique fermentum a quis velit. Aenean eu mollis turpis. Morbi euismod risus eros, at vestibulum lectus iaculis sed. Pellentesque varius eu justo in viverra. Phasellus feugiat tincidunt urna ut accumsan. In ullamcorper vehicula hendrerit. Sed sapien diam, rhoncus nec porta non, gravida et ante. Nulla et volutpat quam, nec venenatis ex. Donec vitae dictum libero, id ornare sapien. Morbi id aliquam ipsum. Cras a ultricies purus.
		
		Morbi sit amet sagittis metus, vitae dictum nunc. Pellentesque ornare consequat diam, vitae maximus dolor facilisis at. Cras semper imperdiet mollis. Fusce maximus augue quis elit pulvinar, vitae varius nunc tempus. Maecenas et suscipit leo. Morbi tempus neque enim, sit amet vestibulum ex sollicitudin ac. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed vel rutrum nisl. Nullam ac vehicula velit. Sed suscipit lacus libero, ac laoreet tortor vehicula a. In aliquam tellus tincidunt ligula fringilla, ut finibus lorem commodo. Sed sit amet porta magna, et aliquam lacus.
		
		Nulla at libero in velit lacinia feugiat sed sit amet erat. Cras nec iaculis magna. Curabitur efficitur turpis vel risus euismod, sit amet posuere orci efficitur. Nunc tincidunt ante eu nunc lacinia, varius dignissim tellus eleifend. Pellentesque enim urna, porttitor eget tincidunt ut, sollicitudin in enim. Donec sit amet hendrerit orci, ut viverra arcu. Nam vitae semper est. Vestibulum aliquet dolor sed turpis blandit eleifend. Nullam auctor accumsan mauris eu mattis. Etiam interdum purus sit amet libero placerat vehicula.
		
		Ut ac odio risus. Sed ut dolor ut metus tincidunt egestas. Curabitur ullamcorper lectus interdum nisi euismod, et faucibus dui fringilla. Nunc at purus vel erat hendrerit dapibus a ut ligula. Nulla facilisi. Etiam gravida, dui nec posuere gravida, nunc quam sodales dui, ut congue mauris tortor vitae nunc. Praesent dapibus mi quis pulvinar consequat. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Pellentesque ultricies placerat dolor tempus cursus. Etiam vehicula, dui vel varius ullamcorper, sapien sem malesuada felis, quis suscipit est nibh vel massa.
		
		Nulla sed mollis enim. Interdum et malesuada fames ac ante ipsum primis in faucibus. Quisque magna purus, faucibus vel ornare vel, semper et odio. Donec non interdum nunc. Maecenas metus augue, maximus ac volutpat id, euismod eget sem. Sed efficitur non diam sed ultrices. Fusce ultrices nisl et suscipit tristique. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi luctus augue vel vehicula iaculis.
		
		Vivamus iaculis luctus nisl, quis molestie purus mollis vitae. Proin vehicula rutrum finibus. Nulla lacinia auctor tincidunt. Vivamus sollicitudin orci eu porttitor vehicula. Vivamus auctor sem sed risus porta cursus. Donec orci dui, lacinia vitae mi quis, tincidunt feugiat elit. Mauris pharetra faucibus justo non volutpat. Praesent pellentesque condimentum quam quis porta. Praesent sed nisi id eros iaculis condimentum nec in libero. Donec eu blandit enim. Aliquam nec elementum libero. Donec nec nisi suscipit felis blandit rutrum. Morbi magna nulla, porttitor id nisl vitae, imperdiet efficitur leo. Ut eu vestibulum justo, vitae dapibus felis.
		
		Quisque ac est nec metus sagittis lobortis mollis id metus. Duis a augue eu erat vulputate facilisis varius ac nunc. Suspendisse eleifend enim suscipit erat aliquet, sed feugiat ligula tempor. Aenean dapibus felis porttitor mauris ultricies vehicula. Vestibulum lacinia scelerisque turpis, ut varius nunc suscipit vitae. Praesent blandit mauris eu semper lacinia. Cras mauris augue, tincidunt quis risus tincidunt, blandit euismod erat. Etiam non sagittis leo, eget pharetra risus.
		
		Donec sodales ultricies neque, non accumsan velit blandit non. Aliquam lacinia orci mauris, commodo porttitor ipsum venenatis sed. Integer sit amet pretium nisi. Ut porttitor, sapien quis gravida egestas, nulla tellus ullamcorper ante, et eleifend diam turpis et odio. Quisque dui orci, commodo sit amet rhoncus ut, pharetra ac leo. Suspendisse non vehicula ex. Morbi gravida lacus vitae ex lacinia, nec aliquam purus consequat. Donec aliquam pretium massa, id viverra nunc blandit sit amet. Donec sed dapibus elit. Cras hendrerit efficitur eros quis malesuada. Aenean a massa in dolor gravida volutpat a et nisi. Nullam sit amet est tempus, condimentum odio egestas, dignissim nibh. Integer tempor id sapien at sagittis.
		
		Aliquam in urna semper, suscipit metus in, pharetra lectus. Proin tempor nibh turpis, a pulvinar justo mollis sit amet. Fusce massa turpis, tristique id semper sit amet, venenatis in mauris. Maecenas id consectetur purus, sed efficitur ipsum. Sed pretium nisi ut sem pulvinar ullamcorper. Nullam at nunc et quam rhoncus egestas eu eu odio. Nullam commodo urna cursus massa porta consectetur. Sed blandit erat ut imperdiet malesuada. Cras nec fringilla ante. Nam sed gravida urna. Nam non dui quis turpis efficitur feugiat. Fusce in quam ex.
		
		Nullam purus libero, egestas eget luctus et, malesuada eget est. Curabitur tristique sollicitudin est, imperdiet cras amet. `), []byte(`password`))
	}
}