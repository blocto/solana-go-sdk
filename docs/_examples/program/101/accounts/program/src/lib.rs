use solana_program::{
    account_info::{next_account_info, AccountInfo},
    entrypoint,
    entrypoint::ProgramResult,
    msg,
    pubkey::Pubkey,
};

entrypoint!(process_instruction);

fn process_instruction(
    _program_id: &Pubkey,
    accounts: &[AccountInfo],
    _instruction_data: &[u8],
) -> ProgramResult {
    let account_info_iter = &mut accounts.iter();

    let first_account_info = next_account_info(account_info_iter)?;
    msg!(&format!(
        "first: {} isSigner: {}, isWritable: {}",
        first_account_info.key.to_string(),
        first_account_info.is_signer,
        first_account_info.is_writable,
    ));

    let second_account_info = next_account_info(account_info_iter)?;
    msg!(&format!(
        "second: {} isSigner: {}, isWritable: {}",
        second_account_info.key.to_string(),
        second_account_info.is_signer,
        second_account_info.is_writable,
    ));

    Ok(())
}
