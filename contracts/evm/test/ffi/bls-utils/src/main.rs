use rand::{self, SeedableRng};

use clap::{Parser, Subcommand};
use serde::{Serialize, Serializer};
use serde_json;
use zeropool_bn::{AffineG1, AffineG2, Fr, Group, G1, G2};

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand, Debug)]
enum Commands {
    Keygen {
        #[arg(short, long)]
        number: u32,

        #[arg(short, long)]
        seed: u32,
    },
    Aggregate {
        priv_keys: Vec<String>,
    },
    AggregateMul {
        #[arg(short, long)]
        multipliers: Vec<String>,

        #[arg(short, long)]
        priv_keys: Vec<String>,
    },
}
struct KeyPair {
    pub_key_g1: AffineG1,
    pub_key_g2: AffineG2,
    priv_key: Fr,
}

#[derive(Serialize)]
struct SolidityKeyPair {
    pub_key_g1: SolidityG1,
    pub_key_g2: SolidityG2,
    priv_key: U256,
}

#[derive(Serialize)]
struct SolidityG1 {
    x: U256,
    y: U256,
}

#[derive(Serialize)]
struct SolidityG2 {
    x: [U256; 2],
    y: [U256; 2],
}

#[derive(Copy, Clone)]
struct U256([u8; 32]);

impl Serialize for U256 {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: Serializer,
    {
        let mut result = Vec::new();

        for byte in self.0 {
            let mut carry = byte as u32;
            for r in result.iter_mut() {
                let x = *r as u32 * 256 + carry;
                *r = (x % 10) as u8;
                carry = x / 10;
            }
            while carry > 0 {
                result.push((carry % 10) as u8);
                carry /= 10;
            }
        }

        if result.is_empty() {
            result.push(0);
        }

        let str: String = result.into_iter().rev().map(|d| (b'0' + d) as char).collect();

        serializer.serialize_str(str.as_str())
    }
}

impl Into<SolidityG1> for AffineG1 {
    fn into(self) -> SolidityG1 {
        let mut g1 = SolidityG1 {
            x: U256([0; 32]),
            y: U256([0; 32]),
        };

        self.x().to_big_endian(&mut g1.x.0).unwrap();
        self.y().to_big_endian(&mut g1.y.0).unwrap();

        g1
    }
}

impl Into<SolidityG2> for AffineG2 {
    fn into(self) -> SolidityG2 {
        let mut g2 = SolidityG2 {
            x: [U256([0; 32]); 2],
            y: [U256([0; 32]); 2],
        };

        self.x().imaginary().to_big_endian(&mut g2.x[0].0).unwrap();
        self.x().real().to_big_endian(&mut g2.x[1].0).unwrap();
        self.y().imaginary().to_big_endian(&mut g2.y[0].0).unwrap();
        self.y().real().to_big_endian(&mut g2.y[1].0).unwrap();

        g2
    }
}

#[derive(Copy, Clone)]
struct CopiablePublicKeyG1(G1);

impl Into<SolidityKeyPair> for KeyPair {
    fn into(self) -> SolidityKeyPair {
        let mut priv_key_bytes = [0; 32];
        self.priv_key.into_u256().to_big_endian(&mut priv_key_bytes).unwrap();

        SolidityKeyPair {
            pub_key_g1: self.pub_key_g1.into(),
            pub_key_g2: self.pub_key_g2.into(),
            priv_key: U256(priv_key_bytes),
        }
    }
}

impl From<Fr> for KeyPair {
    fn from(priv_key: Fr) -> Self {
        let pub_key_g1 = AffineG1::from_jacobian(G1::one() * priv_key).expect("Affine error");
        let pub_key_g2 = AffineG2::from_jacobian(G2::one() * priv_key).expect("Affine error");

        Self {
            priv_key,
            pub_key_g1,
            pub_key_g2,
        }
    }
}

fn aggregate<I, J>(key_pairs: I, multipliers: J) -> KeyPair
where
    I: Iterator<Item = KeyPair>,
    J: Iterator<Item = Fr>,
{
    let (aggregated_priv_key, aggregated_pub_key_g1, aggregated_pub_key_g2) = key_pairs.zip(multipliers).fold(
        (Fr::zero(), G1::zero(), G2::zero()), // Starting point with zero elements for G1 and G2
        |(acc_priv, acc_g1, acc_g2), (key_pair, fr)| {
            (
                acc_priv + key_pair.priv_key * fr,
                acc_g1 + G1::from(key_pair.pub_key_g1) * fr,
                acc_g2 + G2::from(key_pair.pub_key_g2) * fr,
            )
        },
    );

    KeyPair {
        priv_key: aggregated_priv_key,
        pub_key_g1: AffineG1::from_jacobian(aggregated_pub_key_g1).expect("Affine error"),
        pub_key_g2: AffineG2::from_jacobian(aggregated_pub_key_g2).expect("Affine error"),
    }
}

fn main() {
    let cli = Cli::parse();

    match cli.command {
        Some(Commands::Keygen { number, seed }) => {
            let rng = &mut rand::rngs::StdRng::seed_from_u64(seed.into());

            let key_pairs: Vec<KeyPair> = (0..number)
                .map(|_| {
                    let priv_key = Fr::random(rng);

                    KeyPair::from(priv_key)
                })
                .collect();

            let solidity_key_pairs: Vec<SolidityKeyPair> =
                key_pairs.into_iter().map(|keypair| keypair.into()).collect();

            print!("{}", serde_json::to_string(&solidity_key_pairs).unwrap());
        }
        Some(Commands::Aggregate { priv_keys }) => {
            let key_pairs: Vec<KeyPair> = (&priv_keys)
                .into_iter()
                .map(|key| KeyPair::from(Fr::from_str(&key).expect("Not an Fr")))
                .collect();

            let aggregated_key_pair = aggregate(key_pairs.into_iter(), std::iter::repeat(Fr::one()));
            let solidity_aggregated_key_pair: SolidityKeyPair = aggregated_key_pair.into();

            print!("{}", serde_json::to_string(&solidity_aggregated_key_pair).unwrap());
        }
        Some(Commands::AggregateMul { priv_keys, multipliers }) => {
            let key_pairs: Vec<KeyPair> = (&priv_keys)
                .into_iter()
                .map(|key| KeyPair::from(Fr::from_str(&key).expect("Not an Fr")))
                .collect();

            let aggregated_key_pair = aggregate(
                key_pairs.into_iter(),
                multipliers
                    .into_iter()
                    .map(|el| Fr::from_str(el.as_str()).expect("Not an Fr")),
            );
            let solidity_aggregated_key_pair: SolidityKeyPair = aggregated_key_pair.into();

            print!("{}", serde_json::to_string(&solidity_aggregated_key_pair).unwrap());
        }
        None => {}
    }
}
