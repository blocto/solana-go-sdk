use solana_program::{
    account_info::AccountInfo, entrypoint, entrypoint::ProgramResult, msg,
    program_error::ProgramError, pubkey::Pubkey,
};

entrypoint!(process_instruction);

fn process_instruction(
    _program_id: &Pubkey,
    _accounts: &[AccountInfo],
    instruction_data: &[u8],
) -> ProgramResult {
    let (selector, rest) = instruction_data
        .split_first()
        .ok_or(ProgramError::InvalidInstructionData)?;

    match selector {
        0 => msg!(&format!(
            "first instruction is called. remaining data: {:?}",
            rest,
        )),
        1 => msg!(&format!(
            "second instruction is called. remaining data: {:?}",
            rest,
        )),
        _ => {
            msg!("invalid called");
            return Err(ProgramError::InvalidInstructionData);
        }
    }

    Ok(())
}
